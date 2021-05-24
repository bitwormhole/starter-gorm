package drivers

import (
	"github.com/bitwormhole/starter-gorm/data"
	"gorm.io/gorm"
)

type SQLServerDriverAdapter struct {
	data.DBDriverBase
}

func (inst *SQLServerDriverAdapter) Open(params *data.DBOpenParams) (gorm.Dialector, error) {

	//	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	//	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	dsn := "todo..."
	return inst.InvokeOpenFunc(dsn)
}
