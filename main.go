package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
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
			"homePage": "http://localhost:8080",
			"logo":     "/logo.png",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
