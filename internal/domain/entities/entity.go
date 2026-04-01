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
	State  string        `bson:"state"`
	Status string        `bson:"status"`
}

func NewEmployee(name string, state, status string) (*Employee, error) {
	if strings.TrimSpace(state) == "" || strings.TrimSpace(status) == "" {
		return nil, fmt.Errorf("%s", resources.NoneOfTheValuesCanBeEmpty)
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
