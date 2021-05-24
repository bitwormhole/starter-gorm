package main

import (
	"embed"
	"fmt"
	"os"

	etc_demo "github.com/bitwormhole/starter-gorm/demo/etc"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
)

//go:embed src/main/resources
var resources embed.FS

func main() {
	builder := config.NewBuilderFS(&resources, "src/main/resources")

	etc_demo.Config(builder)

	err := tryMain(builder)
	if err != nil {
		panic(err)
	}
}

func tryMain(cb application.ConfigBuilder) error {

	context, err := application.Run(cb.Create(), os.Args)
	if err != nil {
		return err
	}

	err = application.Loop(context)
	if err != nil {
		return err
	}

	code, err := application.Exit(context)
	if err != nil {
		return err
	}

	fmt.Println("exit with code:", code)
	return nil
}
