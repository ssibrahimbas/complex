package value

import (
	"github.com/go-pg/pg/v10"
	"server/databases/postgre"
)

type Value struct {
	Values int `json:"values"`
}

func Init() {
	err := postgre.CreateTable(GetInterface())
	if err != nil {
		panic(err)
	}
}

func GetInterface() interface{} {
	return (*Value)(nil)
}

func GetAll() ([]Value, error) {
	var values []Value
	err := postgre.Query(&values).Select()
	if err != nil {
		return nil, err
	}
	return values, err
}

func Create(val *Value) (pg.Result, error) {
	return postgre.Query(val).Insert()
}
