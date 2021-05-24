package data

import (
	"errors"

	"gorm.io/gorm"
)

type DBOpenFunc func(dsn string) (gorm.Dialector, error)

type DBOpenParams struct {
	SourceName   string // the name of DataSource
	URL          string
	Username     string
	Password     string
	DriverName   string
	DatabaseName string
}

type DBDriver interface {
	Open(params *DBOpenParams) (gorm.Dialector, error)
}

////////////////////////////////////////////////////////////////////////////////

type DBDriverBase struct {
	OpenFunc DBOpenFunc
}

func (inst *DBDriverBase) InvokeOpenFunc(dsn string) (gorm.Dialector, error) {
	fn := inst.OpenFunc
	if fn == nil {
		return nil, errors.New("the.DBOpenFunc==nil")
	}
	return fn(dsn)
}
