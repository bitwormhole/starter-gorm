package datasource

import (
	"errors"
)

////////////////////////////////////////////////////////////////////////////////

type DriverRegistrar struct {
	drivers map[string]Driver
}

func (inst *DriverRegistrar) _try_init() Drivers {
	if inst.drivers == nil {
		inst.drivers = make(map[string]Driver)
	}
	return inst
}

func (inst *DriverRegistrar) Register(name string, driver Driver) {
	inst._try_init()
	if driver == nil {
		return
	}
	inst.drivers[name] = driver
}

func (inst *DriverRegistrar) GetDriverByName(name string) (Driver, error) {
	inst._try_init()
	driver := inst.drivers[name]
	if driver == nil {
		return nil, errors.New("no gorm driver with name:" + name)
	}
	return driver, nil
}

func (inst *DriverRegistrar) Open(cfg *Configuration) (Source, error) {

	if cfg == nil {
		return nil, errors.New("cfg==nil")
	}

	driver, err := inst.GetDriverByName(cfg.Driver)
	if err != nil {
		return nil, err
	}

	return driver.Open(cfg)
}
