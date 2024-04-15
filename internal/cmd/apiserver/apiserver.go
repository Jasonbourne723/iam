package main

import (
	"github.com/jasonbourne723/iam/internal/bootstrap"
	"github.com/jasonbourne723/iam/internal/global"
)

func main() {

	bootstrap.InitializeConfig()

	global.App.DB = bootstrap.InitializeDB()

	bootstrap.RunServe()
}
