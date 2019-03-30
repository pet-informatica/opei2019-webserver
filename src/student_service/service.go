package student_service

import (
	"context"
	"net/http"
	"opei2019-webservice/src/student"
	"opei2019-webservice/src/student_service/store"
)

type Service interface {
	CreateStudent(context.Context, *student.Student) (int32, *student.Student, error)
}

type basicService struct {
	store store.Store
}

func NewService() basicService {
	return basicService{store.New()}
}

func (s basicService) CreateStudent(ctx context.Context, std *student.Student) (int32, *student.Student, error) {
	err := Validate(std)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	student, err := s.store.CreateStudent(std)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, student, nil
}
