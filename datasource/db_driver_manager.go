package datasource

////////////////////////////////////////////////////////////////////////////////

// DriverManager 是驱动管理器
//  [inject:"#starter-gorm-driver-manager"]
type DriverManager interface {
	FindDriver(cfg *Configuration) (Driver, error)
	OpenSource(cfg *Configuration) (Source, error)
}
