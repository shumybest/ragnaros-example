package main

import (
	"github.com/shumybest/ragnaros"
	"ragnaros-example/app"
)

func main() {
	ragnaros.InjectApps(app.DemoController)
	ragnaros.Start()
}
