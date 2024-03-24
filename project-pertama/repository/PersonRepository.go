package repository

import (
	_ "database/sql"
	_ "fmt"
	"project-pertama/model"

	"gorm.io/gorm"
)

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *personRepository{
	return &personRepository{
		db: db,
	}
}

func (pr *personRepository) Create(newPerson model.Person) (model.Person, error){
	// query := "insert into person(name, address) values($1, $2) returning *"

	tx := pr.db.Create(&newPerson)
	return newPerson, tx.Error
}

func (pr *personRepository) GetAll() ([]model.Person, error){
	var persons = []model.Person{}

	tx := pr.db.Find(&persons)
	return persons, tx.Error
}