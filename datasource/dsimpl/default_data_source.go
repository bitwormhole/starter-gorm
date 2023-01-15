package dsimpl

import (
	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"gorm.io/gorm"
)

// DefaultDataSource ... 实现默认的数据源
type DefaultDataSource struct {
	markup.Component `class:"starter-gorm-source-registry life"`

	Drivers datasource.DriverManager `inject:"#starter-gorm-driver-manager"`

	Name            string `inject:"${datasource.default.name}"`
	Driver          string `inject:"${datasource.default.driver}"`
	Username        string `inject:"${datasource.default.username}"`
	Password        string `inject:"${datasource.default.password}"`
	Database        string `inject:"${datasource.default.database}"`
	TableNamePrefix string `inject:"${datasource.default.table-name-prefix}"`
	TableNameSuffix string `inject:"${datasource.default.table-name-suffix}"`
	Host            string `inject:"${datasource.default.host}"`
	Port            int    `inject:"${datasource.default.port}"`
	LazyLoad        bool   `inject:"${datasource.default.lazyload}"`

	target datasource.Source
}

func (inst *DefaultDataSource) _Impl() (datasource.Source, datasource.SourceRegistry, application.LifeRegistry) {
	return inst, inst, inst
}

// GetLifeRegistration ...
func (inst *DefaultDataSource) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnStart: inst.onStart,
	}
}

// ListSources ...
func (inst *DefaultDataSource) ListSources() []*datasource.SourceRegistration {
	name := inst.Name
	reg := &datasource.SourceRegistration{
		Name:   name,
		Source: inst,
	}
	return []*datasource.SourceRegistration{reg}
}

func (inst *DefaultDataSource) onStart() error {
	if !inst.LazyLoad {
		_, err := inst.DB()
		if err != nil {
			return err
		}
	}
	return nil
}

// DB ...
func (inst *DefaultDataSource) DB() (*gorm.DB, error) {
	tar := inst.target
	if tar == nil {
		t2, err := inst.open()
		if err != nil {
			return nil, err
		}
		tar = t2
		inst.target = t2
	}
	return tar.DB()
}

func (inst *DefaultDataSource) open() (datasource.Source, error) {
	cfg := inst.config()
	return inst.Drivers.OpenSource(cfg)
}

func (inst *DefaultDataSource) config() *datasource.Configuration {
	return &datasource.Configuration{
		Name:   inst.Name,
		Driver: inst.Driver,
		Host:   inst.Host,
		Port:   inst.Port,

		Database:        inst.Database,
		TableNamePrefix: inst.TableNamePrefix,
		TableNameSuffix: inst.TableNameSuffix,

		Username: inst.Username,
		Password: inst.Password,
	}
}

// Close ...
func (inst *DefaultDataSource) Close() error {
	return nil // NOP
}
