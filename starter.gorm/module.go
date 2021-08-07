package startergorm

import (
	"github.com/bitwormhole/starter-gorm/etc/gorm"
	"github.com/bitwormhole/starter/application"
)

// Module 方法用于导出 starter-gorm 模块
func Module() application.Module {
	return gorm.ExportModule()
}
