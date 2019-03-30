package student

import "time"

type Student struct {
	ID        string `json:"id"`
	Name      string
	CPF       string `json:"cpf"`
	BirthDate time.Time
	Phone     string
	School    string
	Period    string
	Modality  string
	Level     string
	Census    string
}
