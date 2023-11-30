package db

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

var GQX goqu.DialectWrapper

func InitGoqu() {
	GQX = goqu.Dialect("mysql")
}
