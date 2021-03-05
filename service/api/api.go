package api

import (
	"errors"
	"github.com/aeramu/clean-architecture/entity"
)

type CreateBookReq struct {
}

type CreateBookRes struct {
}

type GetBookReq struct {
	BookID string
}

type GetBookRes struct {
	Book *entity.Book
}

func (req CreateBookReq) Validate() error {
	return nil
}

func (req GetBookReq) Validate() error {
	if req.BookID == "" {
		return errors.New("book id is required")
	}
	return nil
}
