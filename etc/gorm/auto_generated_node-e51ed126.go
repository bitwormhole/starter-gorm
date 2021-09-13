// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gorm

import (
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	markup0x23084a "github.com/bitwormhole/starter/markup"
	gorm0x9f6319 "gorm.io/gorm"
)

type pComConfiguration struct {
	instance *datasource0x68a737.Configuration
	 markup0x23084a.Component `id:"gorm-datasource-config-default"`
	Name string `inject:"${datasource.default.name}"`
	TableNamePrefix string `inject:"${datasource.default.table-name-prefix}"`
	TableNameSuffix string `inject:"${datasource.default.table-name-suffix}"`
	Host string `inject:"${datasource.default.host}"`
	Port int `inject:"${datasource.default.port}"`
	Driver string `inject:"${datasource.default.driver}"`
	Username string `inject:"${datasource.default.username}"`
	Password string `inject:"${datasource.default.password}"`
	Database string `inject:"${datasource.default.database}"`
}


type pComDriverManager struct {
	instance *datasource0x68a737.DriverManager
	 markup0x23084a.Component `id:"gorm-driver-manager"`
	Drivers []datasource0x68a737.Driver `inject:"*"`
}


type pComFacade struct {
	instance *datasource0x68a737.Facade
	 markup0x23084a.Component `id:"gorm-datasource-default" initMethod:"Open" destroyMethod:"Close"`
	Drivers datasource0x68a737.Drivers `inject:"#gorm-driver-manager"`
	Config *datasource0x68a737.Configuration `inject:"#gorm-datasource-config-default"`
	name string ``
	inner datasource0x68a737.Source ``
	db *gorm0x9f6319.DB ``
}

