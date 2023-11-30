package db

import (
	"github.com/bendt-indonesia/env"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// Client is a database client
type Client struct {
	db *sqlx.DB
}

// New instantiates DB
func Init() (*Client, error) {
	DB_HOST := env.Get("DB_HOST")
	DB_USER := env.Get("DB_USER")
	DB_PASS := env.Get("DB_PASS")
	DB_NAME := env.Get("DB_NAME")

	connect := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ")/" + DB_NAME+"?parseTime=true"
	db, err := sqlx.Connect("mysql", connect)
	if err != nil {
		logrus.Panic(err)
	}

	if err = db.Ping(); err != nil {
		logrus.Panic(err)
	}

	logrus.Trace("SQLX Connected!")
	return &Client{db: db}, nil
}
