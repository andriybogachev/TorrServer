package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server/torr"
	"server/web/api"
	"server/web/pages"
)

var (
	BTS      = torr.NewBTS()
	waitChan = make(chan error)
)

func Start(port string) {

	BTS.Connect()

	route := gin.New()
	route.Use(gin.Recovery(), cors.Default())
	api.SetupRoute(route)
	pages.SetupRoute(route)

	waitChan <- route.Run(":" + port)
}

func Wait() error {
	return <-waitChan
}

func Stop() {
	BTS.Disconnect()
	waitChan <- nil
}
