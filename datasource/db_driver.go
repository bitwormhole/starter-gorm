package datasource

const DriversClassName = "gorm-db-drivers"

// Drivers  class='gorm-db-drivers'
type Drivers interface {
	GetDriverByName(name string) (Driver, error)
	Open(cfg *Configuration) (Source, error)
}
