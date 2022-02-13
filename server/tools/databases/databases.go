package databases

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/go-redis/redis/v8"
	"server/config"
	"server/entity"
	"time"
)

type Databases struct {
	config *config.Configuration
}

func NewClient(config *config.Configuration) *Databases {
	return &Databases{
		config: config,
	}
}

func (d Databases) ConnectPostgres() *pg.DB {
	time.Sleep(10 * time.Second)
	db := pg.Connect(&pg.Options{
		User:     d.config.PgUser,
		Password: d.config.PgPassword,
		Addr:     d.config.PgHost + ":" + d.config.PgPort,
		Database: d.config.PgDatabase,
	})
	db.Ping(context.Background())
	err := createPostgresSchema(db)
	if err != nil {
		panic(err)
	}
	return db
}

func (d Databases) ConnectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     d.config.RedisHost + ":" + d.config.RedisPort,
		Password: "",
		DB:       0,
	})
}

func createPostgresSchema(db *pg.DB) error {
	models := []interface{}{
		(*entity.Value)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
