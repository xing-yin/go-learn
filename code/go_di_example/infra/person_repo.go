package infra

import "database/sql"

type PersonRepository struct {
	database *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{database: db}
}

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (r *PersonRepository) FindAll() []*Person {
	rows, err := r.database.Query(`select id,name,age from people;`)
	defer rows.Close()

	if err != nil {
		return nil
	}

	people := []*Person{}

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)

		_ = rows.Scan(&id, &name, &age)

		people = append(people, &Person{
			ID:   id,
			Name: name,
			Age:  age,
		})
	}
	return people
}
