package datasource

import (
	"errors"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
	"gorm.io/gorm"
)

type GormStarterConfig struct {
	DataSourceName string

	//	GormConfig *gorm.Config
	// Injector & Context 二选一必填
	// Injector application.Injector

	Context application.Context
}

type GormDataSource struct {
	dsName  string // the name of DataSource
	drivers Drivers
	inner   Source
	params  *Configuration
	// GormConfig *gorm.Config
	db *gorm.DB
}

func (inst *GormDataSource) _impl_ds_() Source {
	return inst
}

func (inst *GormDataSource) DB() *gorm.DB {
	return inst.inner.DB()
}

func (inst *GormDataSource) Name() string {
	return inst.inner.Name()
}

func (inst *GormDataSource) Configuration() *Configuration {
	return inst.inner.Configuration()
}

func (inst *GormDataSource) loadParams(dsName string, in application.Injector) (*Configuration, error) {

	prefix := "datasource." + dsName + "."
	params := &Configuration{}

	params.Name = dsName
	params.Driver = in.GetPropertyString(prefix+"driver", "[driver_name]")
	params.Database = in.GetPropertyString(prefix+"database", "[db_name]")
	params.Username = in.GetPropertyString(prefix+"username", "[user_name]")
	params.Password = in.GetPropertyString(prefix+"password", "[password]")
	params.Host = in.GetPropertyString(prefix+"host", "[host_name]")
	params.Port = in.GetPropertyInt(prefix+"port", 0)
	params.TableNamePrefix = in.GetPropertyString(prefix+"tablenameprefix", "")
	params.TableNameSuffix = in.GetPropertyString(prefix+"tablenamesuffix", "")

	return params, nil
}

func (inst *GormDataSource) Config(cfg *GormStarterConfig) error {

	if cfg == nil {
		return errors.New("the.GormStarterConfig==nil")
	}

	context := cfg.Context
	if context == nil {
		return errors.New("the.GormStarterConfig.Context==nil")
	}

	injector := context.Injector()
	err := inst.innerConfig(cfg, injector)
	if err != nil {
		return err
	}
	return injector.Done()
}

func (inst *GormDataSource) innerConfig(cfg *GormStarterConfig, in application.Injector) error {

	dataSourceName := cfg.DataSourceName
	params, err := inst.loadParams(dataSourceName, in)
	if err != nil {
		return err
	}

	in.Inject("." + DriversClassName).To(func(o lang.Object) bool {
		drivers, ok := o.(Drivers)
		if ok {
			inst.drivers = drivers
		}
		return ok
	})

	inst.params = params
	inst.dsName = dataSourceName
	return nil
}

func (inst *GormDataSource) Open() error {

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
	driver, err := drivers.GetDriverByName(params.Driver)
	if err != nil {
		return err
	}

	// open inner ds
	source, err := driver.Open(params)
	if err != nil {
		return err
	}

	// 迁移 schema
	// db.AutoMigrate(&Product{})

	inst.db = source.DB()
	inst.inner = source
	return nil
}

func (inst *GormDataSource) Close() error {
	db := inst.db
	inst.db = nil
	if db == nil {
		return nil
	}
	return nil
}
