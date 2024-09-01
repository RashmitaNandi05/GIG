package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/ritankarsaha/backend/controllers"
)

func AuthRoutes(r *gin.Engine) {
    r.POST("/register", controllers.RegisterUser())
    r.POST("/login", controllers.Login())
}