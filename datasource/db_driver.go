package datasource

// Driver 是数据源驱动
type Driver interface {
	Accept(cfg *Configuration) bool
	Open(cfg *Configuration) (Source, error)
}

// DriverRegistration ...
type DriverRegistration struct {
	Driver Driver
}

// DriverRegistry ... [inject:".starter-gorm-driver-registry"]
type DriverRegistry interface {
	GetRegistration() *DriverRegistration
}
