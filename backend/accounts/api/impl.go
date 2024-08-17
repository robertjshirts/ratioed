package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DatabaseInterface interface {
	CreateAccount(email string, username string, bio *string, pfp *string) (*Account, error)
	GetAccounts(username *string) ([]Account, error)
	GetAccountById(accountId string) (*Account, error)
	UpdateAccountById(accountId string, username *string, bio *string, pfp *string) (*Account, error)
	DeleteAccountById(accountId string) error
}

type Server struct {
	db DatabaseInterface
}

func NewServer(database DatabaseInterface) *Server {
	return &Server{
		db: database,
	}
}

func (s *Server) CreateAccount(ctx echo.Context) error {
	var newAccount NewAccount
	if err := ctx.Bind(&newAccount); err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
	}

	fmt.Printf("%v", newAccount)

	account, err := s.db.CreateAccount(newAccount.Email, newAccount.Username, newAccount.Bio, newAccount.Pfp)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("There was an error creating the account: %v", err))
	}

	return ctx.JSON(http.StatusCreated, account)
}

func (s *Server) DeleteAccountById(ctx echo.Context, accountId int) error {
	return ctx.JSON(http.StatusNotImplemented, "DeleteAccountById not impl")
}

func (s *Server) GetAccountById(ctx echo.Context, accountId int) error {
	return ctx.JSON(http.StatusNotImplemented, "GetAccountById not impl")
}

func (s *Server) GetAccounts(ctx echo.Context, params GetAccountsParams) error {
	return ctx.JSON(http.StatusNotImplemented, "GetAccounts not impl")
}

func (s *Server) UpdateAccountById(ctx echo.Context, accountId int) error {
	return ctx.JSON(http.StatusNotImplemented, "UpdateAccountById not impl")
}
