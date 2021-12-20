package startergorm

import (
	"embed"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter-gorm/etc/gorm"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myName     = "github.com/bitwormhole/starter-gorm"
	myVersion  = "v0.0.7"
	myRevision = 7
)

// Module 定义要导出的模块，外部使用请访问【startergorm.Module()】
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myName).Version(myVersion).Revision(myRevision)
	mb.OnMount(gorm.ExportConfig)
	mb.Resources(myResources())
	mb.Dependency(starter.Module())
	return mb.Create()
}

//go:embed src/main/resources
var theResFS embed.FS

func myResources() collection.Resources {
	return collection.LoadEmbedResources(&theResFS, "src/main/resources")
}
