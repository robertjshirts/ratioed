package db

import (
	"fmt"

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

	var post api.Post
	query := `
		INSERT INTO POST (account_id, body)
		VALUES ($1, $2)
		RETURNING id, timestamp
	`
	err = tx.Get(&post, query, accountId, body)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &post, nil
}

func (d *Database) GetPosts(params api.GetPostsParams) ([]api.Post, error) {
	return nil, nil
}

func (d *Database) GetPostById(postId int) (*api.Post, error) {
	return nil, nil
}

func (d *Database) DeletePost(postId int) error {
	return nil
}

func (d *Database) GetIdByUsername(username string) (*int, error) {
	return nil, nil
}
