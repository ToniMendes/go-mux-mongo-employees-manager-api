// Package web provides HTTP routing and handlers for the employees manager API.
package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Routers(router *Handler, port string) error {
	r := mux.NewRouter()

	r.HandleFunc("/api/employees/signup", router.AddNewEmployee).Methods("POST")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return err
	}

	fmt.Printf("server running in port %s", port)
	fmt.Printf("server running in address http://localhost:%s", port)

	return nil
}
