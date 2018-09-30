package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"net/http"
	_ "time"
)

//see example.sh for the use-case

// BuildID is compile-time variable
var BuildID = "0"

//convert your 'main()' into a 'prog(state)'
//'prog()' is run in a child process
func prog(state overseer.State) {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Gin Server ******")
	})

	http.Serve(state.Listener, engine)
}

//then create another 'main' which runs the upgrades
//'main()' is run in the initial process
func main() {
	overseer.Run(overseer.Config{
		Program: prog,
		Address: ":5001",
		Fetcher: &fetcher.File{Path: "gin_server"},
		Debug:   false, //display log of overseer actions
	})
}
