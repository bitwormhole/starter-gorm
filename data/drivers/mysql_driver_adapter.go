package drivers

import (
	"github.com/bitwormhole/starter-gorm/data"
	"gorm.io/gorm"
)

type MySQLDriverAdapter struct {
	data.DBDriverBase
}

func (inst *MySQLDriverAdapter) Open(params *data.DBOpenParams) (gorm.Dialector, error) {

	//	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := ""
	return inst.InvokeOpenFunc(dsn)
}
