package main

import (
	"net/http"

	"github.com/akifkadioglu/youtube-1/database"
	"github.com/akifkadioglu/youtube-1/env"
	"github.com/akifkadioglu/youtube-1/route"
	"github.com/akifkadioglu/youtube-1/utils"
)

func main() {
	env.InitEnv(env.LOCAL)
	utils.InitTokenAuth()
	database.Connection()
	
	r := route.CreateServer()
	http.ListenAndServe(":3000", r)
}
