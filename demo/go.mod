module github.com/bitwormhole/starter-gorm/demo

go 1.16

require (
	github.com/bitwormhole/starter v0.0.0-20210523115635-887cc289d62c // indirect
	github.com/bitwormhole/starter-gorm/datasource v0.0.0 // indirect
	github.com/bitwormhole/starter-gorm/mysql v0.0.0
	github.com/bitwormhole/starter-gorm/sqlserver v0.0.0
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	gorm.io/gorm v1.21.10 // indirect
)

replace github.com/bitwormhole/starter-gorm/datasource => ../datasource

replace github.com/bitwormhole/starter-gorm/mysql => ../mysql

replace github.com/bitwormhole/starter-gorm/sqlserver => ../sqlserver
