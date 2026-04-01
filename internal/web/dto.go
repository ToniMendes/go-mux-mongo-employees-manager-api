package web

import "github.com/go-playground/validator"

type dtoInput struct {
	Name   string `json:"name" validator:"required, min=3, max=60"`
	Email  string `json:"email" validator:"required, email"`
	State  string `json:"state" validator:"required"`
	Status string `json:"status" validator:"required"`
}

func Validate(dto interface{}) error {
	validate := validator.New()

	err := validate.Struct(dto)
	if err != nil {
		return err
	}

	return nil
}
