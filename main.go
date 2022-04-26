package main

import (
	"embed"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os"
)

var envKeys = map[string]string{
	"port": "PORT",
}

//go:embed site/* site/templates/*
var f embed.FS

func getEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New(fmt.Sprintf("%s not set", key))
	}
	return val, nil
}

func BuildRouter() *gin.Engine {
	router := gin.Default()
	router.StaticFS("/site", http.FS(f))
	tmpl := template.Must(template.New("").ParseFS(f, "site/templates/*.tmpl"))
	router.SetHTMLTemplate(tmpl)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"sourceCodeURI": "https://github.com/korenyoni/codefresh-web-app-test",
			"codefreshURI":  "https://g.codefresh.io/",
			"logo":          "/logo.png",
		})
	})

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "alive"})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
			"logo": "/logo.png",
		})
	})

	router.GET("favicon.ico", func(c *gin.Context) {
		file, _ := f.ReadFile("site/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})

	router.GET("logo.png", func(c *gin.Context) {
		file, _ := f.ReadFile("site/logo.png")
		c.Data(
			http.StatusOK,
			"image/png",
			file,
		)
	})

	router.GET("style.css", func(c *gin.Context) {
		file, _ := f.ReadFile("site/style.css")
		c.Data(
			http.StatusOK,
			"text/css",
			file,
		)
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
