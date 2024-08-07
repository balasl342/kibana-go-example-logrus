package routes

import (
	"github.com/balasl342/kibana-go-example-logrus/handlers"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(router *mux.Router, logger *logrus.Logger) {
	router.HandleFunc("/api/students", handlers.ListStudents(logger)).Methods("GET")
	router.HandleFunc("/api/students/{id}", handlers.GetStudent(logger)).Methods("GET")
	router.HandleFunc("/api/students", handlers.AddStudent(logger)).Methods("POST")
	router.HandleFunc("/api/students/{id}", handlers.UpdateStudent(logger)).Methods("PUT")
	router.HandleFunc("/api/students/{id}", handlers.DeleteStudent(logger)).Methods("DELETE")
}
