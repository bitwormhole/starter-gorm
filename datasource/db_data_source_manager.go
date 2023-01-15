package datasource

// SourceManager ... [inject:"#starter-gorm-source-manager"]
type SourceManager interface {
	GetSource(name string) (Source, error)
}
