package main

import (
	_ "mos/boot"
	_ "mos/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
