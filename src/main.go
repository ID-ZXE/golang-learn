package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func main() {
	run()
}

func run() {
	engine := gee.Default()

	engine.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	engine.GET("/panic", func(c *gee.Context) {
		names := []string{"panic"}
		c.String(http.StatusOK, names[100])
	})

	v1 := engine.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello/:name/doc", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		v1.GET("/api/*/doc", func(c *gee.Context) {
			c.String(http.StatusOK, "url %s\n", c.Path)
		})
	}

	v2 := engine.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name/doc", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	engine.Run(":9999")
}

func onlyForV2() gee.HandlerFunc {
	return func(context *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		context.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", context.StatusCode, context.Req.RequestURI, time.Since(t))
	}
}
