package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	router := gin.Default()
	router.Use(TlsHandler())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello world\n")
	})

	// go router.Run(":8085")
	router.RunTLS(":80", "server.crt", "server.key")
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
