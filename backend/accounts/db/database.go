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

func (d *Database) CreateAccount(email string, username string, bio *string, pfp *string) (*api.Account, error) {
	tx, err := d.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Ignore this it's a skill issue i'm not a good dev
	defaultBio := ""
	defaultPfp := ""

	if bio == nil {
		bio = &defaultBio
	}
	if pfp == nil {
		pfp = &defaultPfp
	}

	var id string
	query := `
		INSERT INTO account (email, username, bio, pfp)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	err = tx.QueryRow(query, email, username, *bio, *pfp).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("error creating account: %w", err)
	}

	return d.GetAccountById(id)

}

func (d *Database) GetAccounts(username *string) ([]api.Account, error) {
	return nil, nil
}

func (d *Database) GetAccountById(accountId string) (*api.Account, error) {
	var account api.Account
	query := `
		SELECT * FROM account 
		WHERE id = $1`

	err := d.db.Get(&account, query, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving account: %w", err)
	}

	return &account, nil
}

func (d *Database) UpdateAccountById(accountId string, username *string, bio *string, pfp *string) (*api.Account, error) {
	return nil, nil
}

func (d *Database) DeleteAccountById(accountId string) error {
	return nil
}
