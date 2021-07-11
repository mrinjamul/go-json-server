package main

import (
	"log"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/go-json-server/utils"
)

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	// This will ensure that the angular files are served correctly
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./static/index.html")
		} else {
			c.File("./static/" + path.Join(dir, file))
		}
	})

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

	r.Run()
}

// CORSMiddleware : cross origin resource sharing
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
