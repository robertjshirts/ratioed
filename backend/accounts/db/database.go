package db

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/robertjshirts/ratioed/backend/accounts/api"
	"github.com/robertjshirts/ratioed/backend/accounts/utils"
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

func (d *Database) CreateAccount(email string, username string, bio string, pfp string) (*int, error) {
	return nil, nil
}

func (d *Database) GetAccounts(username *string) ([]api.Account, error) {
	return nil, nil
}

func (d *Database) GetAccountById(accountId int) (*api.Account, error) {
	return nil, nil
}

func (d *Database) UpdateAccountById(accountId int, username *string, bio *string, pfp *string) (*api.Account, error) {
	return nil, nil
}

func (d *Database) DeleteAccountById(accountId int) error {
	return nil
}
