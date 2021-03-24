
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuperchain/dapp-demo/pkg/dapp"
	"log"
)

func main() {
	r := gin.Default()
	r.Static("/index.html","static")

	r.POST("/deploy",dapp.Deploy)
	r.POST("/getLuckId", dapp.GetLuckId)
	r.POST("/getResult", dapp.GetResult)
	r.POST("/startLuckDraw", dapp.StartLuckDraw)
	log.Fatal(r.Run())
}
