package api

type CreateBookReq struct {
}

type CreateBookRes struct {
}

type GetBookReq struct {
}

type GetBookRes struct {
}

func (req CreateBookReq) Validate() error {
	return nil
}

func (req GetBookReq) Validate() error {
	return nil
}
