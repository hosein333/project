// main.go

package main

import (
	"github.com/gin-gonic/gin"

	_ "webapp/docs"
)

var router *gin.Engine

// @title Maktabkhoneh API
// @version 1.0
// @description This is Maktabkhoneh Golang Training API Documentation

// @contact.name API Support
// @contact.url maktabkhoneh.ir
// @contact.email saber.mesgari@gmail.com

// @host 127.0.0.1:8080
// @BasePath
// @query.collection.format multi

func main() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	//router = gin.Default()

	//url := ginSwagger.URL("127.0.0.1:8080/swagger/doc.json") // The url pointing to API definition
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Initialize the routes
	//initRoute()

	// Start serving the application
	//router.Run(":8080")
 r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
