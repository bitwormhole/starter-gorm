package datasource

import (
	"errors"

	"github.com/bitwormhole/starter/markup"
	"gorm.io/gorm"
)

// Facade 数据源的facade对象
type Facade struct {
	markup.Component `id:"gorm-datasource-default" initMethod:"Open" destroyMethod:"Close"`

	Drivers Drivers        `inject:"#gorm-driver-manager"`
	Config  *Configuration `inject:"#gorm-datasource-config-default"`

	//private
	name  string
	inner Source
	db    *gorm.DB
}

func (inst *Facade) _Impl() Source {
	return inst
}

// DB 取DB对象
func (inst *Facade) DB() *gorm.DB {
	return inst.db
}

// Name 取数据源名称
func (inst *Facade) Name() string {
	return inst.name
}

// Configuration 取配置
func (inst *Facade) Configuration() *Configuration {
	return inst.inner.Configuration()
}

// Open 打开数据源
func (inst *Facade) Open() error {

	// for config
	config := inst.Config
	if config == nil {
		return errors.New("datasource.config==nil")
	}

	// get drivers
	drivers := inst.Drivers
	if drivers == nil {
		return errors.New("datasource.drivers==nil")
	}

	// open inner ds
	source, err := drivers.Open(config)
	if err != nil {
		return err
	}

	// 迁移 schema
	// db.AutoMigrate(&Product{})

	inst.db = source.DB()
	inst.name = source.Name()
	inst.inner = source
	return nil
}

// Close 关闭数据源
func (inst *Facade) Close() error {
	inner := inst.inner
	inst.inner = nil
	inst.db = nil
	if inner != nil {
		return inner.Close()
	}
	return nil
}
