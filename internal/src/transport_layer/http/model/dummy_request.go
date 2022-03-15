package model

import (
	"hexagonal/template/internal/src/domain"

	"github.com/go-playground/validator"
)

type DummyRequest struct {
	Id          *string `query:"id"`
	Description *string `query:"description" validate:"required"`
	Limit       int     `query:"limit"`
	Offset      int     `query:"offset"`
}

func (dummyRequest DummyRequest) ToDummy() domain.Dummy {
	return domain.Dummy{
		Id:          dummyRequest.Id,
		Description: dummyRequest.Description,
	}
}

func (dummyRequest DummyRequest) IsValid() error {
	validate := validator.New()
	validateError := validate.Struct(dummyRequest)
	if validateError != nil {
		return validateError
	}
	return nil
}
