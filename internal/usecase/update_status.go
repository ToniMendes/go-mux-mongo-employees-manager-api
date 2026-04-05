package usecase

import (
	"go-mux-mongo-employees-manager/internal/domain"
	"go-mux-mongo-employees-manager/internal/domain/entities"
)

type UpdateUseCaseDtoInput struct {
	Email string
	Status string
}

type UpdateUseCase struct {
	Collection domain.MongoRepository
}

func NewUpdateUseCase(collection domain.MongoRepository) *UpdateUseCase {
	return &UpdateUseCase{
		Collection: collection,
	}
}

func (repo *UpdateUseCase) UpdateStatus(input UpdateUseCaseDtoInput) error{
	onlyStatusModel, err := entities.NewUpdateOnlyStatus(input.Status, input.Email)
	if err != nil {
		return err
	}

	err = repo.Collection.UpdateStatusByEmail(onlyStatusModel)
	if err != nil {
		return err
	}

	return nil
}