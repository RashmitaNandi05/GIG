package middleware

import (
    "net/http"
    "os"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        tokenString := strings.Split(authHeader, " ")[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("SECRET_KEY")), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        userRole := claims["role"].(string)
        if requiredRole != "" && userRole != requiredRole {
            c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
            c.Abort()
            return
        }

        c.Set("user_id", claims["user_id"])
        c.Next()
    }
}