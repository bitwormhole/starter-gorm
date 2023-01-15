package datasource

// Configuration 是数据源的配置
type Configuration struct {

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
