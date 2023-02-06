package core

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"

	"github.com/Imomali1/todo-app/config"
)

type PostgresClient struct {
	DB *sql.DB
}

func InitDB(conf config.Configs) (*PostgresClient, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.DB)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Errorf("Something went wrong with database: %s\n", err.Error())
		return nil, err
	}

	return &PostgresClient{
		DB: db,
	}, nil
}
