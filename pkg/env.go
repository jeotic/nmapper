package pkg

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jinzhu/gorm"
)

type ENV struct {
	DB      *gorm.DB
	Builder goqu.DialectWrapper
}
