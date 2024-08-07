package db

import (
	"errors"
	"sync"

	"github.com/balasl342/kibana-go-example-logrus/models"
)

var (
	students = make(map[string]models.Student)
	mu       sync.Mutex
)

func GetAllStudents() []models.Student {
	mu.Lock()
	defer mu.Unlock()
	var allStudents []models.Student
	for _, student := range students {
		allStudents = append(allStudents, student)
	}
	return allStudents
}

func GetStudentByID(id string) (models.Student, error) {
	mu.Lock()
	defer mu.Unlock()
	student, exists := students[id]
	if !exists {
		return models.Student{}, errors.New("student not found")
	}
	return student, nil
}

func AddStudent(student models.Student) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := students[student.ID]; exists {
		return errors.New("student already exists")
	}
	students[student.ID] = student
	return nil
}

func UpdateStudent(student models.Student) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := students[student.ID]; !exists {
		return errors.New("student not found")
	}
	students[student.ID] = student
	return nil
}

func DeleteStudent(id string) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := students[id]; !exists {
		return errors.New("student not found")
	}
	delete(students, id)
	return nil
}
