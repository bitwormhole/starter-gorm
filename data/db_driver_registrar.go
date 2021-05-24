package data

import (
	"errors"
)

const DBDriversClassName = "gorm-db-drivers"

// DBDrivers class='gorm-db-drivers'
type DBDrivers interface {
	Register(name string, driver DBDriver)
	FindDriverByName(name string) (DBDriver, error)
}

////////////////////////////////////////////////////////////////////////////////

type DBDriverRegistrar struct {
	drivers map[string]DBDriver
}

func (inst *DBDriverRegistrar) _try_init() DBDrivers {
	if inst.drivers == nil {
		inst.drivers = make(map[string]DBDriver)
	}
	return inst
}

func (inst *DBDriverRegistrar) Register(name string, driver DBDriver) {
	inst._try_init()
	if driver == nil {
		return
	}
	inst.drivers[name] = driver
}

func (inst *DBDriverRegistrar) FindDriverByName(name string) (DBDriver, error) {
	inst._try_init()
	driver := inst.drivers[name]
	if driver == nil {
		return nil, errors.New("no gorm driver with name:" + name)
	}
	return driver, nil
}
