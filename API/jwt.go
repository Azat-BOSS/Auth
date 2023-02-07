package API

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JWT_KEY = []byte("SECRET_KEY")

func CreateJWT(userName string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userName,
		"exp": time.Now().Add(time.Hour * 10).Unix(),
	})

	tokenStr, err := token.SignedString(JWT_KEY)
	if err != nil {
		fmt.Println(err)
	}
	return tokenStr
}

func ValidateJWT(next func(c *fiber.Ctx)) {

}
