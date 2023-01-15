package startergorm

import (
	"embed"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter-gorm/gen/cfgstartergorm"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName     = "github.com/bitwormhole/starter-gorm"
	theModuleVersion  = "v0.0.8"
	theModuleRevision = 8
	theModuleResPath  = "src/main/resources"
)

//go:embed src/main/resources
var theModuleResFS embed.FS

// Module 定义要导出的模块，外部使用请访问【startergorm.Module()】
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)

	mb.OnMount(cfgstartergorm.ExportConfig)
	mb.Resources(collection.LoadEmbedResources(&theModuleResFS, theModuleResPath))

	mb.Dependency(starter.Module())
	return mb.Create()
}
