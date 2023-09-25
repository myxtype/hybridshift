package main

import (
	"frame/admin/rest"
	"frame/conf"
)

func main() {
	httpServer := rest.NewHttpServer(conf.GetConfig().AdminServer.Addr)
	httpServer.Start()
}
