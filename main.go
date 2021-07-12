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

	config, err := utils.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	r.Use(CORSMiddleware())

	var (
		webpath string
		port    string
	)

	if config.StaticPath != "" {
		webpath = config.StaticPath
	} else {
		webpath = "static"
	}
	if config.Port != "" {
		port = ":" + config.Port
	} else {
		port = ":3000"
	}

	// This will ensure that the angular files are served correctly
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./" + webpath + "/index.html")
		} else {
			c.File("./" + webpath + "/" + path.Join(dir, file))
		}
	})
	// testing endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	for _, e := range config.Endpoints {
		jdata := utils.GetJSON(e.JSONPath)
		r.GET("/api/"+e.Route, func(c *gin.Context) {
			c.JSON(200, jdata)
		})
	}

	r.Run(port)
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
