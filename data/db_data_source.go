package data

import "gorm.io/gorm"

type DataSource interface {
	Name() string
	DB() *gorm.DB
}
