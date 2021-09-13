package main

import (
	"askUs/v1/package/api"
	"askUs/v1/util/auth"
	"askUs/v1/util/config"
	"askUs/v1/util/db"
	mid "askUs/v1/util/middleware"
	"askUs/v1/util/server"
	"context"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	env := config.Init()
	client := db.Connect(context.Background(), env)
	s := server.Init(env)
	e := s.Start()
	jwtService, err := auth.Init(env)
	handleError(err)
	ap := api.Init(client, jwtService, env, mid.JwtAuth(jwtService))
	ap.Route(e)
	e.Logger.Fatal(e.Start(s.Port))
}

func handleError(e error) {

	if e != nil {
		fmt.Println(e.Error())
	}

}
