package main

import (
	"github.com/bitwormhole/starter"
	startergorm "github.com/bitwormhole/starter-gorm"
)

func main() {
	i := starter.InitApp()
	i.UseMain(startergorm.Module())
	i.Run()
}
