package data

import (
	"errors"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
	"gorm.io/gorm"
)

type GormStarterConfig struct {
	DataSourceName string

	GormConfig *gorm.Config

	// Injector & Context 二选一必填
	Injector application.Injector
	Context  application.Context
}

type GormStarter struct {
	dsName     string // the name of DataSource
	drivers    DBDrivers
	params     *DBOpenParams
	GormConfig *gorm.Config
	db         *gorm.DB
}

func (inst *GormStarter) _impl_ds_() DataSource {
	return inst
}

func (inst *GormStarter) DB() *gorm.DB {
	return inst.db
}

func (inst *GormStarter) Name() string {
	return inst.dsName
}

func (inst *GormStarter) loadParams(dsName string, in application.Injector) (*DBOpenParams, error) {
	params := &DBOpenParams{}
	prefix := "datasource." + dsName + "."

	params.SourceName = dsName
	params.DriverName = in.GetPropertyString(prefix+"driver", "[driver_name]")
	params.DatabaseName = in.GetPropertyString(prefix+"database", "[db_name]")
	params.Username = in.GetPropertyString(prefix+"username", "[user_name]")
	params.Password = in.GetPropertyString(prefix+"password", "[password]")
	params.URL = in.GetPropertyString(prefix+"url", "[db_url]")

	return params, nil
}

func (inst *GormStarter) Config(cfg *GormStarterConfig) error {

	if cfg == nil {
		return errors.New("the.GormStarterConfig==nil")
	}

	injector := cfg.Injector
	if injector != nil {
		return inst.innerConfig(cfg, injector)
	}

	context := cfg.Context
	if context == nil {
		return errors.New("the.GormStarterConfig.Context==nil")
	}

	injector = context.Injector()
	err := inst.innerConfig(cfg, injector)
	if err != nil {
		return err
	}
	return injector.Done()
}

func (inst *GormStarter) innerConfig(cfg *GormStarterConfig, in application.Injector) error {

	dataSourceName := cfg.DataSourceName
	params, err := inst.loadParams(dataSourceName, in)
	if err != nil {
		return err
	}

	in.Inject("." + DBDriversClassName).To(func(o lang.Object) bool {
		drivers, ok := o.(DBDrivers)
		if ok {
			inst.drivers = drivers
		}
		return ok
	})

	inst.GormConfig = cfg.GormConfig
	inst.params = params
	inst.dsName = dataSourceName
	return nil
}

func (inst *GormStarter) Open() error {

	// for params
	params := inst.params
	if params == nil {
		return errors.New("the.DBOpenParams==nil")
	}

	// get drivers
	drivers := inst.drivers
	if drivers == nil {
		return errors.New("drivers==nil")
	}

	// find driver
	driver, err := drivers.FindDriverByName(params.DriverName)
	if err != nil {
		return err
	}

	// open dialector
	dialector, err := driver.Open(params)
	if err != nil {
		return err
	}

	// gorm.Config
	gormConfig := inst.GormConfig
	if gormConfig == nil {
		gormConfig = &gorm.Config{}
	}

	// 打开
	db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		// panic("failed to connect database")
		return err
	}

	// 迁移 schema
	// db.AutoMigrate(&Product{})

	inst.db = db
	return nil
}

func (inst *GormStarter) Close() error {
	db := inst.db
	inst.db = nil
	if db == nil {
		return nil
	}
	return nil
}
