package main

import (
	"github.com/agustfricke/gin-jwt/controllers"
	"github.com/agustfricke/gin-jwt/initialize"
	"github.com/agustfricke/gin-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
    initialize.LoadEnv()
    initialize.DataBase()
}

func main() {
	r := gin.Default()

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.GET("/", middleware.RequireAuth, controllers.Validate)

	r.Run() 
}
