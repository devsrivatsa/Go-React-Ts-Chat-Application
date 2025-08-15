package main

import (
	"log"

	"github.com/devsrivatsa/chat_app_go-ts-react/db"
	"github.com/devsrivatsa/chat_app_go-ts-react/internal/user"
	"github.com/devsrivatsa/chat_app_go-ts-react/internal/ws"
	"github.com/devsrivatsa/chat_app_go-ts-react/router"
)

func main() {
	database, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	userRepository := user.NewRepository(database.GetDB())
	userService := user.NewUserService(userRepository)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHubHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	if err := router.Start("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
