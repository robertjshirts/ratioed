package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type DatabaseInterface interface {
	CreatePost(accountId int, body string, tags []string, attachments []Attachment) (*Post, error)
	GetPosts(username *string, tag *string, sort *GetPostsParamsSort, page *int, limit *int) ([]Post, error)
	GetPostById(postId int) (*Post, error)
	DeletePostById(postId int) error
	GetIdByUsername(username string) (*int, error)
	GetPostCommentsById(postId int) ([]Post, error)
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

	// Parse the hashtags
	tags := getTags(newPost.Content.Body)

	// Get the attachments
	var attachments []Attachment
	if newPost.Content.Attachments != nil {
		attachments = *newPost.Content.Attachments
	}

	// Get the username's account id
	id, err := s.db.GetIdByUsername(newPost.Username)

	// If err when trying to get account
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprint(err))
	}

	// If no account found
	if id == nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("There is no user with username '%s'", newPost.Username))
	}

	// Create the post
	post, err := s.db.CreatePost(*id, newPost.Content.Body, tags, attachments)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("error creating post: %v", err))
	}

	return ctx.JSON(http.StatusCreated, *post)
}

// DeletePostById handles the DELETE /posts/{postId} request.
func (s *Server) DeletePostById(ctx echo.Context, postId string) error {
	// TODO: auth
	// try to convert postId to an int
	id, err := strconv.Atoi(postId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "postId in path parameter must be an integer")
	}

	err = s.db.DeletePostById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("error deleting post: %v", err))
	}

	return ctx.NoContent(http.StatusNoContent)
}

// GetPostById handles the GET /posts/{postId} request.
func (s *Server) GetPostById(ctx echo.Context, postId string) error {
	// try to convert postId to an int
	id, err := strconv.Atoi(postId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "postId in path parameter must be an integer")
	}

	post, err := s.db.GetPostById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("error retrieving post: %v", err))
	}

	if post == nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, post)
}

func (s *Server) GetPostCommentsById(ctx echo.Context, postId string) error {
	// try to convert postId to an int
	id, err := strconv.Atoi(postId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "postId in path parameter must be an integer")
	}

	posts, err := s.db.GetPostCommentsById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("error retrieving posts: %s", err))
	}

	return ctx.JSON(http.StatusOK, posts)
}

func getTags(body string) []string {
	var tags []string
	var buffer strings.Builder
	inTag := false
	for _, char := range body {
		if char == '#' {
			if inTag == true {
				if buffer.Len() > 0 {
					tags = append(tags, buffer.String())
				}
				buffer.Reset()
				continue
			}
			inTag = true

			// Continue so we don't push the hash to the buffer
			continue
		}

		if char == ' ' {
			if inTag {
				inTag = !inTag
				tags = append(tags, buffer.String())
				buffer.Reset()
			}

			// Continue so we don't push the space to the buffer
			continue
		}

		if inTag {
			buffer.WriteRune(char)
		}
	}

	if buffer.Len() > 0 {
		tags = append(tags, buffer.String())
	}

	return tags
}
