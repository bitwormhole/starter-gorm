// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfgstartergorm

import (
	datasource0x68a737 "github.com/bitwormhole/starter-gorm/datasource"
	dsimpl0x33f250 "github.com/bitwormhole/starter-gorm/datasource/dsimpl"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComDefaultDataSource struct {
	instance *dsimpl0x33f250.DefaultDataSource
	 markup0x23084a.Component `class:"starter-gorm-source-registry life"`
	Drivers datasource0x68a737.DriverManager `inject:"#starter-gorm-driver-manager"`
	Name string `inject:"${datasource.default.name}"`
	Driver string `inject:"${datasource.default.driver}"`
	Username string `inject:"${datasource.default.username}"`
	Password string `inject:"${datasource.default.password}"`
	Database string `inject:"${datasource.default.database}"`
	TableNamePrefix string `inject:"${datasource.default.table-name-prefix}"`
	TableNameSuffix string `inject:"${datasource.default.table-name-suffix}"`
	Host string `inject:"${datasource.default.host}"`
	Port int `inject:"${datasource.default.port}"`
	LazyLoad bool `inject:"${datasource.default.lazyload}"`
}


type pComDriverManagerImpl struct {
	instance *dsimpl0x33f250.DriverManagerImpl
	 markup0x23084a.Component `id:"starter-gorm-driver-manager"`
	Drivers []datasource0x68a737.DriverRegistry `inject:".starter-gorm-driver-registry"`
}


type pComSourceManagerImpl struct {
	instance *dsimpl0x33f250.SourceManagerImpl
	 markup0x23084a.Component `id:"starter-gorm-source-manager"`
	Sources []datasource0x68a737.SourceRegistry `inject:".starter-gorm-source-registry"`
}

