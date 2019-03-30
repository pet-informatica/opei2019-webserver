package student_service

import "opei2019-webservice/src/student"

type ValidationError struct{}

func (e ValidationError) Error() string {
	return "Student not valid"
}

func Validate(std *student.Student) error {
	// Create all validations for student here
	if std.CPF == "" {
		return ValidationError{}
	}

	return nil
}
