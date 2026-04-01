// Package usecase contains the business logic use cases for the employees manager API.
package usecase

import (
	"go-mux-mongo-employees-manager/internal/domain"
	"go-mux-mongo-employees-manager/internal/domain/entities"
)

type UseCaseDtoInput struct {
	Name   string
	State  string
	Status string
}

type CreateUseCase struct {
	Collection domain.MongoRepository
}

func NewCreateUseCase(collection domain.MongoRepository) *CreateUseCase {
	return &CreateUseCase{
		Collection: collection,
	}
}

func (repo *CreateUseCase) Create(dto UseCaseDtoInput) error {
	model, err := entities.NewEmployee(dto.Name, dto.State, dto.Status)
	if err != nil {
		return err
	}

	err = repo.Collection.Insert(model)
	if err != nil {
		return err
	}

	return nil
}
