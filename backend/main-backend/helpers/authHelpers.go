package helpers

import (
    "os"
    "time"
    "github.com/golang-jwt/jwt/v4"
	"errors"
	"log"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWT(userID string, role string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
func CheckUserType(c *gin.Context, roll string) error {
	userType := c.GetString("user_type")
	if userType != roll {
		return errors.New("user type of the user cannot be matched by the database")
	}
	return nil
}

func MathUserTypeToUid(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	id := c.GetString("uid")

	if userType == "USER" && id != userId {
		return errors.New("user id or uid cannot be properly matched for the user")
	}

	err := CheckUserType(c, userType)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("Error hashing password:", err)
		return ""
	}
	return string(bytes)
}