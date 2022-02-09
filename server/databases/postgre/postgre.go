package postgre

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB

func Connect(options *pg.Options) {
	db = pg.Connect(options)
	defer db.Close()
}

func CreateTable(model interface{}) error {
	return db.Model(model).CreateTable(&orm.CreateTableOptions{
		Temp: true,
	})
}

func Query(model interface{}) *pg.Query {
	return db.Model(model)
}
