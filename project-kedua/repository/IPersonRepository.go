package repository

import "project-kedua/model"

type IPersonRepository interface {
	Create(newPerson model.Person) (model.Person, error)
	GetAll() ([]model.Person, error)
	Update(id int, newPerson model.Person) (model.Person, error)
	Delete(uuid string) error
}