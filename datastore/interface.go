package datastore

import (
	"database/sql"
	"golayered/entities"
)

type Animal interface {
	Get(id int) ([]entities.Animal, error)
	Create(entities.Animal) (entities.Animal, error)
}

type AnimalStore struct {
	db *sql.DB
}

func NewAnimalStore(db *sql.DB) AnimalStore {
	return AnimalStore{db: db}
}

func (a AnimalStore) Get(id int) ([]entities.Animal, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if id != 0 {
		rows, err = a.db.Query("SELECT * from animals where id = ?", id)
	} else {
		rows, err = a.db.Query("SELECT  * FROM animals")
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var animals []entities.Animal

	for rows.Next() {
		var a entities.Animal
		_ = rows.Scan(&a.ID, &a.Name, &a.Age)
	}
	return animals, nil
}

func (a AnimalStore) Create(animal entities.Animal) (entities.Animal, error) {
	res, err := a.db.Exec("INSERT INTO animals (name, age) "+
		"VALUES(?,?)", animal.Name, animal.Age)

	if err != nil {
		return entities.Animal{}, nil
	}

	id, _ := res.LastInsertId()
	animal.ID = int(id)
	return animal, nil
}
