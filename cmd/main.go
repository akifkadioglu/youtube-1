package main

import (
	"fmt"
	"net/http"

	"github.com/akifkadioglu/youtube-1/env"
	"github.com/akifkadioglu/youtube-1/models"
	"github.com/akifkadioglu/youtube-1/route"
	"github.com/akifkadioglu/youtube-1/utils"
)

func main() {
	env.InitEnv(env.LOCAL)
	fmt.Println(env.Getenv(env.PASSWORD))
	
	utils.InitTokenAuth()
	fmt.Println(utils.GenerateToken(models.JwtModel{
		Name:"akif",
		Age:21,
	}))
	r := route.CreateServer()
	http.ListenAndServe(":3000", r)
}
