package main

import (
	"fmt"
	"net/http"
	"prueba/handlers"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Servidor corriendo en el puerto 8080")
	http.ListenAndServe(":8080", mux)
}
