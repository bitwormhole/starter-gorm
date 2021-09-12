package gorm

import "github.com/bitwormhole/starter/application"

// ExportConfig 导出配置
func ExportConfig(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
