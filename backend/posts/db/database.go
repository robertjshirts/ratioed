package db

import (
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

func (d *Database) CreatePost(accountId int, body string, tags []string, attachments []api.Attachment) (*api.Post, error) {
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

	var id int
	query := `
		INSERT INTO post (account_id, body)
		VALUES ($1, $2)
		RETURNING id`
	err = tx.QueryRow(query, accountId, body).Scan(&id)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error creating post in database: %w", err)
	}

	// TODO: Insert attachments

	// TODO: Insert tags

	// Return the post details

	tx.Commit()
	return d.GetPostById(id)
}

func (d *Database) GetPosts(username *string, tag *string, sort *api.GetPostsParamsSort, limit *int, page *int) ([]api.Post, error) {
	var tempPosts []tempPost
	// Do left joins because of potential query params
	query := `
		SELECT p.id, p.parent_id, p.body, p.account_id, p.ratioed, p.timestamp
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

	err := d.db.Select(&tempPosts, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error retrieving posts: %w", err)
	}

	var posts []api.Post

	for _, tempPost := range tempPosts {
		posts = append(posts, *tempPost.ToPost())
	}

	return posts, nil
}

func (d *Database) GetPostById(postId int) (*api.Post, error) {
	var post tempPost

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

	// Query for the post
	query := `
	        SELECT id, parent_id, body, account_id, ratioed, timestamp
	        FROM post
	        WHERE id = $1`
	err = tx.Get(&post, query, postId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving post from db: %w", err)
	}

	// Query for the attachments
	// Ignore error bc it'll only ever be no rows
	query = `
        SELECT id, src
        FROM attachment
        WHERE parent_id = $1`
	tx.Select(&post.Content.Attachments, query, postId)

	post.Likes = 0
	post.Dislikes = 0

	// Get the actual username
	post.Username, err = d.GetUsernameById(post.AccountId)
	if err != nil {
		return nil, fmt.Errorf("error getting username: %w", err)
	}

	return post.ToPost(), nil
}

func (d *Database) DeletePost(postId int) error {
	return nil
}

func (d *Database) GetUsernameById(account_id int) (string, error) {
	var username string

	query := `
		SELECT username 
		FROM account
		WHERE id = $1`
	if err := d.db.Get(&username, query, account_id); err != nil {
		return "", err
	}

	return username, nil
}

func (d *Database) GetIdByUsername(username string) (int, error) {
	id := 10
	return id, nil
}

func (d *Database) GetPostCommentsById(postId int) ([]api.Post, error) {
	var tempComments []tempPost
	query := `
	        SELECT id, parent_id, body, account_id, ratioed, timestamp
	        FROM post
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
