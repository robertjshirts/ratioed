package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DatabaseInterface interface {
	CreatePost(accountId int, body string, tags []string, attachments []Attachment) (*Post, error)
	GetPosts(username *string, tag *string, sort *GetPostsParamsSort, page *int, limit *int) ([]Post, error)
	GetPostById(postId int) (*Post, error)
	DeletePost(postId int) error
	GetIdByUsername(username string) (int, error)
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
	posts, err := s.db.GetPosts(params.Username, params.Tag, params.Sort, params.Limit, params.Page)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("There was an error getting the posts: %v", err))
	}
	return ctx.JSON(http.StatusOK, posts)

}

// CreatePost handles the POST /posts request.
func (s *Server) CreatePost(ctx echo.Context) error {
	// Get the json
	var newPost NewPost
	if err := ctx.Bind(&newPost); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Get the username's account id
	id, _ := s.db.GetIdByUsername(newPost.Username)

	// Create the post
	post, err := s.db.CreatePost(id, newPost.Content.Body, nil, nil)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("error creating post: %v", err))
	}

	return ctx.JSON(http.StatusCreated, post)
}

// DeletePostById handles the DELETE /posts/{postId} request.
func (s *Server) DeletePostById(ctx echo.Context, postId string) error {
	// TODO: Implement logic to delete the post by ID
	// Example: return ctx.NoContent(http.StatusNoContent)
	return ctx.JSON(http.StatusNotImplemented, "DeletePostById not implemented")
}

// GetPostById handles the GET /posts/{postId} request.
func (s *Server) GetPostById(ctx echo.Context, postId string) error {
	// try to convert postId to an int
	id, err := strconv.Atoi(postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "postId in path parameter must be an integer")
	}

	post, err := s.db.GetPostById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("error retrieving post: %v", err))
	}

	return ctx.JSON(http.StatusOK, post)
}
