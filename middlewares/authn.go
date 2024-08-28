package middlewares

import (
	"fmt"
	"os"
	"rooming_house/entities"
	"rooming_house/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(userID uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return utils.HandlerError(c, utils.NewUnauthorizedError("please login first"))
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return utils.HandlerError(c, utils.NewUnauthorizedError("invalid authorization header format"))
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			return utils.HandlerError(c, utils.NewUnauthorizedError("invalid token"))
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["user_id"].(float64))
			username := claims["username"].(string)

			c.Set("userPayload", &entities.JWTUserPayload{
				UserID:   userID,
				Username: username,
			})
		} else {
			return utils.HandlerError(c, utils.NewUnauthorizedError("invalid token"))
		}

		return next(c)
	}
}
