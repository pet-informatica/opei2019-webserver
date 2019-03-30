package postgres

import (
	"opei2019-webservice/src/student"

	"github.com/go-pg/pg"
)

var connection *pg.DB

type BasicDatabase interface {
	GetConnection() *pg.DB
	CloseConnection()
}

type StudentDatabase interface {
	BasicDatabase

	CreateStudent(*student.Student) (*student.Student, error)
}

type basicDatabase struct {
}

type studentDatabase struct {
	BasicDatabase
}

func NewDatabase() studentDatabase {
	var database studentDatabase
	basicDatabase := basicDatabase{}

	database = studentDatabase{
		basicDatabase,
	}

	return database
}

func (d basicDatabase) CloseConnection() {
	if connection != nil {
		connection.Close()
	}
}

func (d basicDatabase) GetConnection() *pg.DB {
	if connection == nil {
		addr := "localhost"
		port := "1234"
		user := "admin"
		pass := "admin"
		name := "OPEI_DB"

		connection = pg.Connect(&pg.Options{
			User:     user,
			Password: pass,
			Database: name,
			Addr:     addr + ":" + port,
			PoolSize: 30,
		})
	}

	return connection
}

func (d studentDatabase) CreateStudent(std *student.Student) (*student.Student, error) {
	db := d.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	err = tx.Insert(&std)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return std, nil
}
