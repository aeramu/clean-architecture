package api

import "github.com/aeramu/clean-architecture/entity"

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
	return nil
}
