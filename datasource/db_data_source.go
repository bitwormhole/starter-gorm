package datasource

import "gorm.io/gorm"

type Configuration struct {
	// URL      string
	Name     string // the name of DataSource
	Host     string
	Port     int
	Driver   string
	Username string
	Password string
	Database string
}

type Source interface {
	Name() string
	DB() *gorm.DB
	Configuration() *Configuration
}

type Driver interface {
	Open(cfg *Configuration) (Source, error)
}
