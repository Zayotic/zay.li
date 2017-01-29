package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/style.css", "./static/css/style.css")
  r.StaticFile("/image.gif", "./static/images/image.gif")
  r.StaticFile("/music.mp3", "./static/music.mp3")
  r.StaticFile("/music2.mp3", "./static/music2.mp3")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/ssh", func(c *gin.Context) {
		c.HTML(http.StatusOK, "key.html", gin.H{})
	})

	// Custom Server Info
	s := &http.Server{
		Addr:           ":80",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
