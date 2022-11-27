package main

import (
	"locallife/pkg"
	"locallife/pkg/log"
	"locallife/router"
)

func main() {
	pkg.Init()
	r := router.SetRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("server start failed.", log.ErrorF(err))
	}
}
