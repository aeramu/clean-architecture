package api

import (
	"errors"
	"github.com/aeramu/clean-architecture/entity"
	"strings"
)

type CreateBookReq struct {
	Title string
	Author string
	CoverImage string
}

type CreateBookRes struct {
	Message string
}

type GetBookReq struct {
	BookID string
}

type GetBookRes struct {
	Book *entity.Book
}

func (req CreateBookReq) Validate() error {
	if req.Title == "" {
		return errors.New("title is required")
	}
	if req.Author == "" {
		return errors.New("author is required")
	}
	if req.CoverImage == "" {
		return errors.New("cover image is required")
	}
	if strings.Contains(req.CoverImage, " ") {
		return errors.New("cover image cannot contain space")
	}
	return nil
}

func (req GetBookReq) Validate() error {
	if req.BookID == "" {
		return errors.New("book id is required")
	}
	return nil
}
