package golang

import (
	"testing"

	startergorm "github.com/bitwormhole/starter-gorm"
	"github.com/bitwormhole/starter/tests"
)

func TestGorm(t *testing.T) {

	i := tests.Starter(t)
	i.Use(startergorm.Module())

	rt := i.RunTest()

	rt.Exit()
}
