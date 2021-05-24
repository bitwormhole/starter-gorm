package drivers

import (
	"github.com/bitwormhole/starter-gorm/data"
	"gorm.io/gorm"
)

type PostgresDriverAdapter struct {
	data.DBDriverBase
}

func (inst *PostgresDriverAdapter) Open(params *data.DBOpenParams) (gorm.Dialector, error) {

	//	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dsn := ""
	return inst.InvokeOpenFunc(dsn)
}
