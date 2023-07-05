package main

import (
	"github.com/agustfricke/gin-jwt/initialize"
	"github.com/agustfricke/gin-jwt/models"
)


func init() {
    initialize.LoadEnv()
    initialize.DataBase()
}

func main() {
    initialize.DB.AutoMigrate(&models.User{})
}

