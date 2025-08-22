package router

import (
	"time"

	"github.com/devsrivatsa/chat_app_go-ts-react/internal/user"
	"github.com/devsrivatsa/chat_app_go-ts-react/internal/ws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.HubHandler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "http://10.0.0.88:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	auth := r.Group("/")
	auth.Use(AuthRequired())
	auth.POST("/ws/createRoom", wsHandler.CreateRoom)
	auth.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom) //ws://localhost:8080/ws/joinRoom/1?clientId=1&username=first
	auth.GET("/ws/getRooms", wsHandler.GetRooms)
	auth.GET("/ws/getClients/:roomId", wsHandler.GetClientsInRoom)
}

func Start(address string) error {
	return r.Run(address)
}
