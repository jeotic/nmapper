package pkg

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
)

type ENV struct {
	DB      *sql.DB
	Builder goqu.DialectWrapper
}
