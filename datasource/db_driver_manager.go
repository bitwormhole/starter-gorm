package datasource

import (
	"errors"

	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// DriverManager 是驱动管理器
type DriverManager struct {
	markup.Component `id:"gorm-driver-manager"`
	Drivers          []Driver `inject:".gorm-datasource-driver"`
}

func (inst *DriverManager) _Impl() Drivers {
	return inst
}

func (inst *DriverManager) list() []Driver {
	list := inst.Drivers
	if list == nil {
		list = make([]Driver, 0)
	}
	return list
}

// FindDriver 根据配置查找驱动
func (inst *DriverManager) FindDriver(cfg *Configuration) (Driver, error) {
	if cfg == nil {
		return nil, errors.New("cfg==nil")
	}
	all := inst.list()
	for _, driver := range all {
		if driver.Accept(cfg) {
			return driver, nil
		}
	}
	return nil, errors.New("no gorm driver with name:" + cfg.Driver)
}

// Open 打开数据源
func (inst *DriverManager) Open(cfg *Configuration) (Source, error) {

	if cfg == nil {
		return nil, errors.New("cfg==nil")
	}

	driver, err := inst.FindDriver(cfg)
	if err != nil {
		return nil, err
	}

	return driver.Open(cfg)
}
