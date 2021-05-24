package drivers

import (
	"github.com/bitwormhole/starter-gorm/data"
	"gorm.io/gorm"
)

type SQLiteDriverAdapter struct {
	data.DBDriverBase
}

func (inst *SQLiteDriverAdapter) Open(params *data.DBOpenParams) (gorm.Dialector, error) {

	// github.com/mattn/go-sqlite3
	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	dsn := ""
	return inst.InvokeOpenFunc(dsn)
}
