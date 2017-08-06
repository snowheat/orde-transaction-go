package main

import (
	"github.com/buaazp/fasthttprouter"
	OTRouter "github.com/snowheat/order-transaction-go/application/v1/router"
	OTServer "github.com/snowheat/order-transaction-go/application/v1/system/server"
	"github.com/valyala/fasthttp"
)

func main() {

	router := fasthttprouter.New()

	server := fasthttp.Server{}

	OTRouter.Set(router)
	OTServer.Set(&server, router)
	OTServer.Run(&server)
}
