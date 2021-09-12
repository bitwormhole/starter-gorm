package datasource

import "gorm.io/gorm"

////////////////////////////////////////////////////////////////////////////////

// GormDataSourceBuilder 数据源创建者
type GormDataSourceBuilder struct {
	// DSN       string
	Dialector gorm.Dialector
	Config1   Configuration
	Config2   gorm.Config
}

// Open 打开数据源
func (inst *GormDataSourceBuilder) Open() (Source, error) {
	src := &innerGormDataSource{}
	return src.open(inst)
}

////////////////////////////////////////////////////////////////////////////////

type innerGormDataSource struct {
	db   *gorm.DB
	conf Configuration
	name string
}

func (inst *innerGormDataSource) open(builder *GormDataSourceBuilder) (Source, error) {

	db, err := gorm.Open(builder.Dialector, &builder.Config2)
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
	dst := &Configuration{}
	*dst = inst.conf
	return dst
}

func (inst *innerGormDataSource) Name() string {
	return inst.name
}

func (inst *innerGormDataSource) Close() error {
	inst.db = nil
	return nil
}

////////////////////////////////////////////////////////////////////////////////
