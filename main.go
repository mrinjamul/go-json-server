package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/go-json-server/utils"
)

func main() {
	r := gin.Default()

	config, err := utils.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	for _, e := range config.Endpoints {
		jdata := utils.GetJSON(e.JSONPath)
		r.GET(e.Route, func(c *gin.Context) {
			c.JSON(200, jdata)
		})
	}

	r.Run(":3000")
}
