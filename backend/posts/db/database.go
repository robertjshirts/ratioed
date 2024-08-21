package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/robertjshirts/ratioed/backend/posts/api"
	"github.com/robertjshirts/ratioed/backend/posts/utils"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase() *Database {
	dbHost := utils.GetEnv("DB_HOST")
	dbPort := utils.GetEnv("DB_PORT")
	dbUser := utils.GetEnv("DB_USER")
	dbPassword := utils.GetEnv("DB_PASSWORD")
	dbName := utils.GetEnv("DB_NAME")

	connStr := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName
	fmt.Println("Connecting to database: ", connStr)

	db := sqlx.MustConnect("postgres", connStr)

	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	return &Database{db: db}
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) CreatePost(accountId string, parentId *string, body string, tags []string, attachments []api.Attachment) (*api.Post, error) {
	tx, err := d.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var id string
	query := `
		INSERT INTO post (account_id, parent_id, body)
		VALUES ($1, $2, $3)
		RETURNING id`

	if parentId != nil {
		err = tx.QueryRow(query, accountId, parentId, body).Scan(&id)
	} else {
		err = tx.QueryRow(query, accountId, nil, body).Scan(&id)
	}
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error creating post in database: %w", err)
	}

	if len(attachments) > 0 {
		query = `
			INSERT INTO attachment (parent_id, src, type)
			VALUES `
		var attachmentValues []string
		for _, attachment := range attachments {
			attachmentValues = append(attachmentValues, fmt.Sprintf("(%d, '%s', '%s')", id, attachment.Src, attachment.Type))
		}
		query += strings.Join(attachmentValues, ", ")
		_, err := tx.Exec(query)
		if err != nil {
			return nil, fmt.Errorf("error creating attachments: %w", err)
		}
	}
	if len(tags) > 0 {
		query = `
			INSERT INTO hashtag (parent_id, tag)
			VALUES `
		var tagValues []string
		for _, tag := range tags {
			tagValues = append(tagValues, fmt.Sprintf("(%d, '%s')", id, tag))
		}
		query += strings.Join(tagValues, ", ")
		_, err := tx.Exec(query)
		if err != nil {
			return nil, fmt.Errorf("error creating tags, %w", err)
		}
	}

	// Return the post details

	tx.Commit()
	return d.GetPostById(id)
}

func (d *Database) GetPosts(username *string, tag *string, sort *api.GetPostsParamsSort, limit *int, page *int) ([]api.Post, error) {
	tx, err := d.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var tempPosts []tempPost
	// Do left joins because of potential query params
	query := `
		SELECT p.id, p.parent_id, p.body, p.account_id, p.ratioed, p.timestamp, a.username
		FROM post p
		LEFT JOIN account a on p.account_id = a.id
		LEFT JOIN hashtag h on p.id = h.parent_id`
	args := []interface{}{}
	conditions := []string{}
	argIdx := 1

	if username != nil {
		conditions = append(conditions, fmt.Sprintf("a.username = $%d", argIdx))
		args = append(args, *username)
		argIdx++
	}

	if tag != nil {
		conditions = append(conditions, fmt.Sprintf("h.tag = $%d", argIdx))
		args = append(args, *tag)
		argIdx++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Sort order
	if sort != nil && (*sort == "asc") {
		query += " ORDER BY p.timestamp ASC"
	} else {
		query += " ORDER BY p.timestamp DESC"
	}

	// Set limit
	query += fmt.Sprintf(" LIMIT $%d", argIdx)
	if limit == nil {
		args = append(args, 100)
		argIdx++
		if page != nil {
			query += fmt.Sprintf(" OFFSET $%d", argIdx)
			args = append(args, ((*page - 1) * 100))
		}
	} else {
		args = append(args, *limit)
		argIdx++
		if page != nil {
			query += fmt.Sprintf(" OFFSET $%d", argIdx)
			args = append(args, ((*page - 1) * *limit))
		}
	}

	err = tx.Select(&tempPosts, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error retrieving posts: %w", err)
	}

	// Convert tempPost to reg Post + get attachments
	var posts []api.Post
	for _, tempPost := range tempPosts {
		attachments, err := d.GetPostAttachmentsById(*tx, tempPost.Id)
		if err != nil {
			return nil, fmt.Errorf("error getting attachments for posts: %w", err)
		}
		tempPost.Attachments = &attachments
		posts = append(posts, *tempPost.ToPost())
	}

	return posts, nil
}

func (d *Database) GetPostById(postId string) (*api.Post, error) {
	// Start the transaction
	tx, err := d.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("error starting transaction with db: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var post tempPost

	// Query for the post
	query := `
	        SELECT p.id, p.parent_id, p.body, p.account_id, p.ratioed, p.timestamp, a.username
	        FROM post p
		LEFT JOIN account a on p.account_id = a.id
		WHERE p.id = $1`
	err = tx.Get(&post, query, postId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving post from db: %w", err)
	}

	attachments, err := d.GetPostAttachmentsById(*tx, postId)
	if err != nil {
		return nil, fmt.Errorf("error getting attachments: %w", err)
	}

	post.Attachments = &attachments

	// Impl likes at a later date (probably requires another db :(((
	post.Likes = 0
	post.Dislikes = 0

	return post.ToPost(), nil
}

func (d *Database) DeletePostById(postId string) error {
	query := `
		DELETE from post
		WHERE id = $1`
	_, err := d.db.Exec(query, postId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return fmt.Errorf("error deleting post: %v", err)
	}

	return nil
}

func (d *Database) GetUsernameById(account_id string) (*string, error) {
	var username string

	query := `
		SELECT username 
		FROM account
		WHERE id = $1`
	err := d.db.Get(&username, query, account_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &username, nil
}

func (d *Database) GetIdByUsername(username string) (*string, error) {
	var id string
	query := `
		SELECT id
		FROM account
		WHERE account.username = $1`
	err := d.db.Get(&id, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting account id from username: %w", err)
	}
	return &id, nil
}

func (d *Database) GetPostCommentsById(postId string) ([]api.Post, error) {
	var tempComments []tempPost
	query := `
	        SELECT p.id, p.parent_id, p.body, p.account_id, p.ratioed, p.timestamp, a.username
	        FROM post p
		LEFT JOIN account a on p.account_id = a.id
		WHERE parent_id = $1`
	err := d.db.Select(&tempComments, query, postId)
	if err != nil {
		return nil, fmt.Errorf("error getting comments: %w", err)
	}

	var comments []api.Post
	for _, comment := range tempComments {
		comments = append(comments, *comment.ToPost())
	}

	return comments, nil
}

func (d *Database) GetPostAttachmentsById(tx sqlx.Tx, postId string) ([]api.Attachment, error) {
	var attachments []api.Attachment
	query := `
		SELECT src, type FROM attachment
		WHERE attachment.parent_id = $1`
	err := tx.Select(&attachments, query, postId)
	if err != nil {
		return nil, fmt.Errorf("error getting post attachments: %w", err)
	}

	return attachments, nil
}
