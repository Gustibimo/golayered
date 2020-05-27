package datastore

import (
	"database/sql"
	"golayered/entities"
)

type Animal interface {
	Get(id int) ([]entities.Animal, error)
	Create(entities.Animal) (entities.Animal, error)
}
