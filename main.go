package main

import (
    "go-first/database"
    "go-first/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    database.Connect()

    router := gin.Default()
    routes.SetupRoutes(router)
    router.Run(":8080")
}