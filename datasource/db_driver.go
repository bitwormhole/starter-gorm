package datasource

const DriversClassName = "gorm-db-drivers"

// Drivers  class='gorm-db-drivers'
type Drivers interface {
	FindDriver(cfg *Configuration) (Driver, error)
	Open(cfg *Configuration) (Source, error)
}
