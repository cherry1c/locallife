package main

import (
	"locallife/pkg"
	"locallife/pkg/conf"
	"locallife/pkg/log"
	"locallife/router"
	"strconv"
)

func main() {
	pkg.Init()
	r := router.SetRouter()
	ip := conf.GetString(conf.ProjectName + ".server.ip")
	port := conf.GetInt(conf.ProjectName + ".server.http.port")
	if err := r.Run(ip + ":" + strconv.Itoa(port)); err != nil {
		log.Fatal("server start failed.", log.ErrorF(err))
	}
}
