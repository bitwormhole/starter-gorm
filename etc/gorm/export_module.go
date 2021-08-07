package gorm

import (
	"github.com/bitwormhole/starter-gorm/etc/gorm/conf"
	"github.com/bitwormhole/starter/application"
)

// ExportModule 定义要导出的模块，外部使用请访问【startergorm.Module()】
func ExportModule() application.Module {
	return &application.DefineModule{
		Name:     "github.com/bitwormhole/starter-gorm",
		Version:  "v0.0.0",
		Revision: 1,
		OnMount:  func(cb application.ConfigBuilder) error { return cfg(cb) },
	}
}

func cfg(cb application.ConfigBuilder) error {
	return conf.Config(cb)
}
