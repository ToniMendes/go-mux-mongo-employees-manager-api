// Package repositories contains repository interfaces for data access operations.
package repositories

import "go-mux-mongo-employees-manager/internal/usecase"

type IWriterOnlyRepository interface {
	Create(usecase.UseCaseDtoInput) error
	UpdateStatus(usecase.UpdateUseCaseDtoInput) error
}
