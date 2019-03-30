package store

import (
	"opei2019-webservice/src/student"
	"opei2019-webservice/src/student_service/store/postgres"
)

type Store interface {
	CreateStudent(*student.Student) (*student.Student, error)
}

type basicStore struct {
}

func New() basicStore {
	return basicStore{}
}

func (s basicStore) CreateStudent(std *student.Student) (*student.Student, error) {
	database := postgres.NewDatabase()
	return database.CreateStudent(std)
}
