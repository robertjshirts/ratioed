package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DatabaseInterface interface {
	CreateAccount(email string, username string, bio string, pfp string) (*int, error)
	GetAccounts(username *string) ([]Account, error)
	GetAccountById(accountId int) (*Account, error)
	UpdateAccountById(accountId int, username *string, bio *string, pfp *string) (*Account, error)
	DeleteAccountById(accountId int) error
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
	return ctx.JSON(http.StatusNotImplemented, "CreateAccount not impl")
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
