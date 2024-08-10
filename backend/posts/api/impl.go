package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DatabaseInterface interface {
	CreatePost(accountId int, body string, tags []string, attachments []Attachment) (*Post, error)
	GetPosts(params GetPostsParams) ([]Post, error)
	GetPostById(postId int) (*Post, error)
	DeletePost(postId int) error
	GetIdByUsername(username string) (*int, error)
}

type Server struct {
	db DatabaseInterface
}

func NewServer(database DatabaseInterface) *Server {
	return &Server{
		db: database,
	}
}

// GetPosts handles the GET /posts request.
func (s *Server) GetPosts(ctx echo.Context, params GetPostsParams) error {
	// TODO: Implement logic to retrieve posts based on params
	return ctx.JSON(http.StatusNotImplemented, "GetPosts not implemented")
}

// CreatePost handles the POST /posts request.
func (s *Server) CreatePost(ctx echo.Context) error {
	var newPost NewPost
	if err := ctx.Bind(&newPost); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	return ctx.JSON(http.StatusNotImplemented, "not impl")
}

// DeletePostById handles the DELETE /posts/{postId} request.
func (s *Server) DeletePostById(ctx echo.Context, postId string) error {
	// TODO: Implement logic to delete the post by ID
	// Example: return ctx.NoContent(http.StatusNoContent)
	return ctx.JSON(http.StatusNotImplemented, "DeletePostById not implemented")
}

// GetPostById handles the GET /posts/{postId} request.
func (s *Server) GetPostById(ctx echo.Context, postId string) error {
	// TODO: Implement logic to retrieve a specific post by ID
	// Example: return ctx.JSON(http.StatusOK, post)
	return ctx.JSON(http.StatusNotImplemented, "GetPostById not implemented")
}
