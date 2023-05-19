package main

import (
	"baby-chain/gpp"
	"baby-chain/gpp/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func publicRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/public", func(ctx *gin.Context) {
		gpp.SaveHyperParams()
	})
	clientRoute.GET("/sync", services.SyncGet)
	clientRoute.POST("/sync", services.SyncPost)
	clientRoute.POST("/register_ins", services.RegisterIns)
	clientRoute.GET("/list_ins", services.INSList)
}

func privateRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/private", func(ctx *gin.Context) {
		gpp.SaveHyperParams()
	})
	clientRoute.POST("/node", services.NodePost)
	clientRoute.POST("/buy_ins", services.BuyIns)
	clientRoute.POST("/claim_ins", services.ClaimIns)
	clientRoute.POST("/status", services.Status)
}

func main() {
	gpp.BC, gpp.SD, gpp.CSAlgo = gpp.FetchHyperParams()
	gpp.SaveHyperParams()
	chainName := "baby_chain"
	go func() {
		server := gin.Default()
		server.Use(cors.Default())
		basePath := server.Group("/" + chainName)
		privateRoutes(basePath)
		log.Fatalln(server.Run(":9080"))
	}()
	server := gin.Default()
	server.Use(cors.Default())
	basePath := server.Group("/" + chainName)
	publicRoutes(basePath)
	log.Fatalln(server.Run(":9090"))
}
