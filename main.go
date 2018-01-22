package main

import (
	psmconfig "github.com/plagiari-sm/psm-config"
	"github.com/plagiari-sm/svc-serve-static/app"
)

func init() {
	psmconfig.NewConfig()
}

func main() {
	app := new(app.APP)
	app.Initialize()
	app.Serve()
}
