package datasource

import "gorm.io/gorm"

////////////////////////////////////////////////////////////////////////////////

type GormDataSourceBuilder struct {
	// DSN       string
	Dialector gorm.Dialector
	Config1   *Configuration
	Config2   *gorm.Config
}

func (inst *GormDataSourceBuilder) Create() (Source, error) {
	src := &innerGormDataSource{}
	return src.init(inst)
}

////////////////////////////////////////////////////////////////////////////////

type innerGormDataSource struct {
	db   *gorm.DB
	conf *Configuration
	name string
}

func (inst *innerGormDataSource) init(builder *GormDataSourceBuilder) (Source, error) {

	db, err := gorm.Open(builder.Dialector, builder.Config2)
	if err != nil {
		return nil, err
	}

	inst.name = builder.Config1.Name
	inst.conf = builder.Config1
	inst.db = db
	return inst, nil
}

func (inst *innerGormDataSource) DB() *gorm.DB {
	return inst.db
}

func (inst *innerGormDataSource) Configuration() *Configuration {
	src := inst.conf
	dst := &Configuration{}
	*dst = *src
	return dst
}

func (inst *innerGormDataSource) Name() string {
	return inst.name
}

////////////////////////////////////////////////////////////////////////////////
