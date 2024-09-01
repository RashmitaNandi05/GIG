package main

import (
    "log"
    "net/http"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    // "github.com/ritankarsaha/backend/controllers"
    "github.com/ritankarsaha/backend/database"
    "github.com/ritankarsaha/backend/routes"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    r := gin.Default()

    database.ConnectDB()
    routes.AuthRoutes(r)
    routes.UserRoutes(r)
    routes.JobRoutes(r)

    log.Fatal(http.ListenAndServe(":8080", r))
}