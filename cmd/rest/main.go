package main

import (
	"frame/conf"
	"frame/rest"
)

func main() {
	httpServer := rest.NewHttpServer(conf.GetConfig().RestServer.Addr)
	httpServer.Start()
}
