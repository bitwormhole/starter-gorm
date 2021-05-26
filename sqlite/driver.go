package sqlite

import (
	"errors"

	"github.com/bitwormhole/starter-gorm/datasource"
	driver_sqlserver "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Driver struct {
}

func (inst *Driver) Open(params *datasource.Configuration) (datasource.Source, error) {

	//	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	//	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	dsn := "todo..."

	dialector := driver_sqlserver.Open(dsn)
	if dialector == nil {
		return nil, errors.New("driver_sqlserver.Open() return nil")
	}

	gc := &gorm.Config{}

	builder := &datasource.GormDataSourceBuilder{}
	builder.Config1 = params
	builder.Config2 = gc
	builder.Dialector = dialector
	return builder.Create()
}
