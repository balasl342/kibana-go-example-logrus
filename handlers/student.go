package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/balasl342/kibana-go-example-logrus/db"
	"github.com/balasl342/kibana-go-example-logrus/models"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func ListStudents(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.WithFields(logrus.Fields{
			"method":    r.Method,
			"endpoint":  r.URL.Path,
			"timestamp": time.Now(),
		}).Info("List students endpoint hit")

		students := db.GetAllStudents()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	}
}

func GetStudent(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		logger.WithFields(logrus.Fields{
			"method":    r.Method,
			"endpoint":  r.URL.Path,
			"id":        id,
			"timestamp": time.Now(),
		}).Info("Get student endpoint hit")

		student, err := db.GetStudentByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(student)
	}
}

func AddStudent(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student models.Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logger.WithFields(logrus.Fields{
			"method":    r.Method,
			"endpoint":  r.URL.Path,
			"student":   student,
			"timestamp": time.Now(),
		}).Info("Add student endpoint hit")

		if err := db.AddStudent(student); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateStudent(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		var student models.Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		student.ID = id

		logger.WithFields(logrus.Fields{
			"method":    r.Method,
			"endpoint":  r.URL.Path,
			"student":   student,
			"timestamp": time.Now(),
		}).Info("Update student endpoint hit")

		if err := db.UpdateStudent(student); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteStudent(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		logger.WithFields(logrus.Fields{
			"method":    r.Method,
			"endpoint":  r.URL.Path,
			"id":        id,
			"timestamp": time.Now(),
		}).Info("Delete student endpoint hit")

		if err := db.DeleteStudent(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
