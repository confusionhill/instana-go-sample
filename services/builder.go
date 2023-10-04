package services

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func LoadBuilder() goqu.DialectWrapper {
	dialect := goqu.Dialect("postgres")
	return dialect
}
