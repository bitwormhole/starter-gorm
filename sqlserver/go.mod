module github.com/bitwormhole/starter-gorm/sqlserver

go 1.16

require (
	github.com/bitwormhole/starter v0.0.0-20210523115635-887cc289d62c // indirect
	github.com/bitwormhole/starter-gorm/datasource v0.0.0
	gorm.io/driver/sqlserver v1.0.7 // indirect
	gorm.io/gorm v1.21.10 // indirect
)

replace github.com/bitwormhole/starter-gorm/datasource => ../datasource
