package etc

import (
	"errors"

	"github.com/bitwormhole/starter-gorm/data"
	"github.com/bitwormhole/starter-gorm/data/drivers"
	"github.com/bitwormhole/starter/application"
	"gorm.io/gorm"
)

func _gorm_starter(inst *data.GormStarter, ctx application.Context) error {

	//[component]
	//  initMethod=Open
	//  destroyMethod=Close

	err := inst.Config(&data.GormStarterConfig{
		DataSourceName: "demo",
		Context:        ctx,
	})
	return err
}

func _gorm_db_drivers(inst *data.DBDriverRegistrar, ctx application.Context) error {

	//[component]
	//  class= gorm-db-drivers

	mysql_driver := &drivers.MySQLDriverAdapter{}
	mysql_driver.OpenFunc = func(dsn string) (gorm.Dialector, error) {
		return nil, errors.New("this is a mock driver for mysql")
	}
	inst.Register("mysql", mysql_driver)

	return nil
}
