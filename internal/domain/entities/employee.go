// Package entities contains the domain entities for the employees manager API.
package entities

import (
	"fmt"
	"go-mux-mongo-employees-manager/internal/domain/entities/enum"
	"go-mux-mongo-employees-manager/internal/resources"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Employee struct {
	ID     bson.ObjectID `bson:"_id,omitempty"`
	Name   string        `bson:"name"`
	Email  string        `bson:"email"`
	State  string        `bson:"state"`
	Status string        `bson:"status"`
}

type UpdateOnlyStatus struct {
	Email  string `bson:"email"`
	Status string `bson:"status"`
}

func NewEmployee(name, email, state, status string) (*Employee, error) {
	if strings.TrimSpace(state) == "" || strings.TrimSpace(status) == "" {
		return nil, fmt.Errorf("%s", resources.NoneOfTheValuesCanBeEmpty)
	}

	if strings.TrimSpace(email) == "" {
		return nil, fmt.Errorf("%s", resources.EmailCannotBeEmpty)
	}

	if !strings.Contains(email, "@") || !strings.Contains(email, ".com") {
		return nil, fmt.Errorf("%s", resources.EmailInvalid)
	}

	name, err := formatName(name)
	if err != nil {
		return nil, err
	}

	if status != "0" && status != "1" {
		return nil, fmt.Errorf("%s", resources.StatusInvalid)
	}

	if status == "1" {
		status = enum.StatusActive.String()
	} else {
		status = enum.StatusInactive.String()
	}

	return &Employee{
		Name:   name,
		Email:  email,
		State:  state,
		Status: status,
	}, nil
}

func formatName(name string) (string, error) {
	nameFormated := strings.TrimSpace(name)

	if nameFormated == "" {
		return "", fmt.Errorf("%s", resources.NoneOfTheValuesCanBeEmpty)
	}

	caser := cases.Title(language.BrazilianPortuguese)

	return caser.String(nameFormated), nil
}

func NewUpdateOnlyStatus(status, email string) (*UpdateOnlyStatus, error) {
	if strings.TrimSpace(email) == "" {
		return nil, fmt.Errorf("%s", resources.EmailCannotBeEmpty)
	}

	if !strings.Contains(email, "@") || !strings.Contains(email, ".com") {
		return nil, fmt.Errorf("%s", resources.EmailInvalid)
	}

	if status != "0" && status != "1" {
		return nil, fmt.Errorf("%s", resources.StatusInvalid)
	}

	if status == "1" {
		status = enum.StatusActive.String()
	} else {
		status = enum.StatusInactive.String()
	}

	return &UpdateOnlyStatus{
		Status: status,
		Email:  email,
	}, nil
}
