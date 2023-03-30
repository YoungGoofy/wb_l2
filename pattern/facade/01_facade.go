package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Паттерн Фасад относится к классу структурных паттернов, он нужен для упрощения работы со сложными системы путем создания простого интерфейса
доступа к функционалу этой системы
*/

type Database struct {
	data map[int]string
}

func (d *Database) Get(id int) (string, error) {
	if val, ok := d.data[id]; ok {
		return val, nil
	}
	return "", fmt.Errorf("error")
}

func (d *Database) Set(id int, val string) {
	d.data[id] = val
}

type DatabaseFasad struct {
	db *Database
}

func NewDatabaseFasade() *DatabaseFasad {
	return &DatabaseFasad{
		db: &Database{
			data: make(map[int]string),
		},
	}
}

func (df *DatabaseFasad) Get(id int) (string, error) {
	return df.db.Get(id)
}

func (df *DatabaseFasad) Set(id int, val string) {
	df.db.Set(id, val)
}

func main() {
	fasade := NewDatabaseFasade()
	fasade.Set(1, "hello")
	fmt.Println(fasade.Get(1))
	fmt.Println(fasade.Get(2))
}
