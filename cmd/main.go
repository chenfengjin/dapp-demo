
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuperchain/dapp-demo/pkg/luck_draw"
	"log"
)

func main() {
	r := gin.Default()
	r.Static("/index.html","static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/deploy",luck_draw.Deploy)
	r.POST("/getLuckId", luck_draw.GetLuckId)
	r.POST("/getResult",luck_draw.GetResult)
	r.POST("/startLuckDraw",luck_draw.StartLuckDraw)
	log.Fatal(r.Run())
}
