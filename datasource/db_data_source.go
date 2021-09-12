package datasource

import (
	"io"

	"github.com/bitwormhole/starter/markup"
	"gorm.io/gorm"
)

// Configuration 是数据源的配置
type Configuration struct {
	markup.Component `id:"gorm-datasource-config-default"`

	// URL      string
	Name            string `inject:"${datasource.default.name}"` // the name of DataSource
	TableNamePrefix string `inject:"${datasource.default.table-name-prefix}"`
	TableNameSuffix string `inject:"${datasource.default.table-name-suffix}"`

	Host     string `inject:"${datasource.default.host}"`
	Port     int    `inject:"${datasource.default.port}"`
	Driver   string `inject:"${datasource.default.driver}"`
	Username string `inject:"${datasource.default.username}"`
	Password string `inject:"${datasource.default.password}"`
	Database string `inject:"${datasource.default.database}"`
}

// Source 是数据源
type Source interface {
	io.Closer

	Name() string
	DB() *gorm.DB
	Configuration() *Configuration
}

// Driver 是数据源驱动
type Driver interface {
	Accept(cfg *Configuration) bool
	Open(cfg *Configuration) (Source, error)
}
