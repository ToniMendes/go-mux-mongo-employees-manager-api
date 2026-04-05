// Package domain contains the domain models and interfaces for the employees manager API.
package domain

import "go-mux-mongo-employees-manager/internal/domain/entities"

type MongoRepository interface {
	Insert(*entities.Employee) error
	UpdateStatusByEmail(*entities.UpdateOnlyStatus) error
}
