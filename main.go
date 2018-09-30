package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"net/http"
)

func prog(state overseer.State) {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Gin Server")
	})

	http.Serve(state.Listener, engine)
}

func main() {
	overseer.Run(overseer.Config{
		Program: prog,
		Address: ":5001",
		Fetcher: &fetcher.File{Path: "gin_server"},
		Debug:   false,
	})
}
