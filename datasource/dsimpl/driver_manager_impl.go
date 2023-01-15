package dsimpl

import (
	"errors"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
)

// DriverManagerImpl ....
type DriverManagerImpl struct {
	markup.Component `id:"starter-gorm-driver-manager"`

	Drivers []datasource.DriverRegistry `inject:".starter-gorm-driver-registry"`
}

func (inst *DriverManagerImpl) _Impl() datasource.DriverManager {
	return inst
}

func (inst *DriverManagerImpl) listAllDrivers() []datasource.Driver {
	src := inst.Drivers
	dst := make([]datasource.Driver, 0)
	for _, reg1 := range src {
		if reg1 == nil {
			continue
		}
		reg2 := reg1.GetRegistration()
		if reg2 == nil {
			continue
		}
		drv := reg2.Driver
		if drv == nil {
			continue
		}
		dst = append(dst, drv)
	}
	return dst
}

// FindDriver 根据配置查找驱动
func (inst *DriverManagerImpl) FindDriver(cfg *datasource.Configuration) (datasource.Driver, error) {

	if cfg == nil {
		return nil, errors.New("cfg==nil")
	}

	all := inst.listAllDrivers()
	for _, driver := range all {
		if driver.Accept(cfg) {
			return driver, nil
		}
	}

	return nil, errors.New("no gorm driver with name:" + cfg.Driver)
}

// OpenSource 打开数据源
func (inst *DriverManagerImpl) OpenSource(cfg *datasource.Configuration) (datasource.Source, error) {

	if cfg == nil {
		return nil, errors.New("cfg==nil")
	}

	driver, err := inst.FindDriver(cfg)
	if err != nil {
		return nil, err
	}

	return driver.Open(cfg)
}
