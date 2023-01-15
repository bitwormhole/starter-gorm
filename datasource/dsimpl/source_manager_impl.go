package dsimpl

import (
	"fmt"

	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/markup"
)

// SourceManagerImpl ....
type SourceManagerImpl struct {
	markup.Component `id:"starter-gorm-source-manager"`

	Sources []datasource.SourceRegistry `inject:".starter-gorm-source-registry"`
}

func (inst *SourceManagerImpl) _Impl() datasource.SourceManager {
	return inst
}

// GetSource ...
func (inst *SourceManagerImpl) GetSource(name string) (datasource.Source, error) {
	src := inst.Sources
	for _, reg1 := range src {
		list := reg1.ListSources()
		for _, reg2 := range list {
			if reg2.Name == name && reg2.Source != nil {
				return reg2.Source, nil
			}
		}
	}
	return nil, fmt.Errorf("no data source with name:" + name)
}
