package web

import (
	"encoding/json"
	"fmt"
	"go-mux-mongo-employees-manager/internal/usecase"
	"go-mux-mongo-employees-manager/internal/usecase/repositories"
	"net/http"
	"sync"
)

type UseCaseRepositories interface {
	repositories.IWriterOnlyRepository
}

type Handler struct {
	usecase UseCaseRepositories
}

func NewHandler(uc UseCaseRepositories) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) AddNewEmployee(w http.ResponseWriter, r *http.Request) {
	var dto []dtoInput
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type employeeError struct {
		Err   error
		Index int
	}

	totalEmployees := len(dto)
	var wg sync.WaitGroup
	var errChan = make(chan employeeError, len(dto))

	for i, employee := range dto {

		wg.Add(1)

		go func(e dtoInput, index int) {
			defer wg.Done()

			err := Validate(e)
			if err != nil {
				errChan <- employeeError{Index: index, Err: err}
				return
			}

			dtoUseCase := usecase.UseCaseDtoInput{
				Name:   e.Name,
				Email:  e.Email,
				State:  e.State,
				Status: e.Status,
			}

			err = h.usecase.Create(dtoUseCase)
			if err != nil {
				errChan <- employeeError{Index: index, Err: err}
				return
			}

		}(employee, i)
	}
	wg.Wait()
	close(errChan)

	w.Header().Set("Content-Type", "application/json")

	var errList []string
	for c := range errChan {
		resp := fmt.Sprintf("%s, employee [%d]", c.Err.Error(), c.Index)
		errList = append(errList, resp)
	}

	if len(errList) > 0 {
		if len(errList) == totalEmployees {
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "None of the employees were created",
				"errors":  errList,
			})

			return
		} else {
			w.WriteHeader(http.StatusMultiStatus)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "some employees were not created",
				"errors":  errList,
			})
			return
		}
	}

	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(map[string]string{"message": "employee created successfully"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
}

func (h *Handler) UpdateStatusEmployee(w http.ResponseWriter, r *http.Request) {
	var dto dtoInput

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := Validate(dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dtoUseCase := usecase.UpdateUseCaseDtoInput{
		Email:  dto.Email,
		Status: dto.Status,
	}

	err := h.usecase.UpdateStatus(dtoUseCase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "employee status updated successfully"}); err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
