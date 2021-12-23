package main

import (
	"github.com/gin-gonic/gin"
	"nkn-node-manager/controllers"
	"nkn-node-manager/models"
)

func main() {

	r := gin.Default()

	models.ConnectDatabase() // new

	r.POST("/wallet", controllers.UploadWallet)
	r.POST("/walletform", controllers.UploadWalletFile)
	r.GET("/wallet/:address", controllers.FindWallet)
	r.GET("/generateid/:address", controllers.FindWallet)
	r.GET("/wallet/", controllers.FindIdleWallet)

	r.Run()
}