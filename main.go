package main

import (
	"github.com/aristotelis79/Jsonzable/controllers"
	"github.com/savsgio/atreugo/v10"
	"github.com/valyala/fasthttp"
)

func main() {

	server := configServer()

	registerRoutes(server)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func configServer() *atreugo.Atreugo {
	config := atreugo.Config{Addr: "127.0.0.1:8000"}
	server := atreugo.New(config)
	return server
}

func registerRoutes(server *atreugo.Atreugo) {
	baseControllerPath := "/base"
	server.Path(fasthttp.MethodGet, baseControllerPath, controllers.GET)
	server.Path(fasthttp.MethodPost, baseControllerPath, controllers.POST)
	server.Path(fasthttp.MethodPut, baseControllerPath, controllers.PUT)
	server.Path(fasthttp.MethodDelete, baseControllerPath, controllers.DELETE)
}
