package web

import (
	"encoding/json"
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
	var wg sync.WaitGroup

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, employee := range dto {
		wg.Add(1)

		go func() {
			defer wg.Done()

			err := Validate(employee)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			dtoUseCase := usecase.UseCaseDtoInput{
				Name:   employee.Name,
				Email:  employee.Email,
				State:  employee.State,
				Status: employee.Status,
			}

			err = h.usecase.Create(dtoUseCase)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}()
	}
	wg.Wait()


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(map[string]string{"message": "employee created successfully"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
}
