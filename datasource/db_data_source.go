package datasource

import (
	"io"

	"gorm.io/gorm"
)

// Source 是数据源
type Source interface {
	io.Closer

	// Name() string
	// Configuration() *Configuration

	DB() (*gorm.DB, error)
}

// SourceRegistration ...
type SourceRegistration struct {
	Name   string
	Source Source
}

// SourceRegistry ... [inject:".starter-gorm-source-registry"]
type SourceRegistry interface {
	ListSources() []*SourceRegistration
}
