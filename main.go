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

func getEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New(fmt.Sprintf("%s not set", key))
	}
	return val, nil
}

func BuildRouter() *gin.Engine {
	router := gin.Default()
	router.StaticFile("/favicon.ico", "./site/favicon.ico")
	router.StaticFile("/style.css", "./site/style.css")
	router.StaticFile("/logo.png", "./site/logo.png")
	router.LoadHTMLGlob("site/templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"sourceCodeURI": "https://github.com/korenyoni/codefresh-web-app-test",
			"codefreshURI":  "https://g.codefresh.io/",
			"logo":          "/logo.png",
		})
	})
	router.GET("/health-check", func(c *gin.Context) {
		c.String(http.StatusOK, "alive")
	})
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
			"logo": "/logo.png",
		})
	})
	return router
}

func main() {
	port, err := getEnv(envKeys["port"])
	if err != nil {
		log.Fatal(err)
	}

	router := BuildRouter()
	err = router.Run(fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatal(err)
	}
}
