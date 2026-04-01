package web

import (
	"encoding/json"
	"go-mux-mongo-employees-manager/internal/usecase"
	"go-mux-mongo-employees-manager/internal/usecase/repositories"
	"net/http"
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
	var dto dtoInput

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := Validate(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dtoUseCase := usecase.UseCaseDtoInput{
		Name:   dto.Name,
		State:  dto.State,
		Status: dto.Status,
	}

	err = h.usecase.Create(dtoUseCase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(map[string]string{"message": "employee created successfully"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
}
