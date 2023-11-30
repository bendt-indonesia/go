package db

import (
	"github.com/bendt-indonesia/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var DBX *sqlx.DB

func InitSqlX() {
	DB_HOST := env.Get("DB_HOST")
	DB_USER := env.Get("DB_USER")
	DB_PASS := env.Get("DB_PASS")
	DB_NAME := env.Get("DB_NAME")

	connect := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ")/" + DB_NAME + "?parseTime=true"
	dbc, err := sqlx.Connect("mysql", connect)
	if err != nil {
		logrus.Panic(err)
	}

	if err = dbc.Ping(); err != nil {
		logrus.Panic(err)
	}

	DBX = dbc
	logrus.Trace("SQLX Connected!")
}
