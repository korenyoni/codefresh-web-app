package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var envKeys = map[string]string{
	"port": "PORT",
}

func main() {
	port, err := getEnv(envKeys["port"])
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.StaticFile("/favicon.ico", "./site/favicon.ico")
	r.StaticFile("/style.css", "./site/style.css")
	r.StaticFile("/logo.png", "./site/logo.png")
	r.LoadHTMLGlob("site/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"sourceCodeURI": "https://github.com/korenyoni/codefresh-web-app-test",
			"codefreshURI":  "https://g.codefresh.io/",
			"logo":          "/logo.png",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.tmpl", gin.H{
			"logo": "/logo.png",
		})
	})
	r.Run(fmt.Sprintf(":%s", port))
}

func getEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New(fmt.Sprintf("%s not set", key))
	}
	return val, nil
}
