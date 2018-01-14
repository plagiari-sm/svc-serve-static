package main

import (
	"github.com/plagiari-sm/svc-serve-static/app"
	"github.com/plagiari-sm/svc-serve-static/conf"
)

func init() {
	conf.NewConf()
}

func main() {
	app := new(app.APP)
	app.Initialize()
	app.Serve()
}
