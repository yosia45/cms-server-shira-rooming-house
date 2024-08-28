package main

import (
	"log"
	"rooming_house/config"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	// echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()

	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = "8080"
	// }

	e := echo.New()

	// e.GET("/swagger/*", echoSwagger.WrapHandler)
	// e.GET("/activities", handler.GetAllActivities, middleware.JWTAuth)

	// cli.RegisterUserRoutes(e)
	// cli.RegisterPostRoutes(e)
	// cli.RegisterCommentRoutes(e)

	// e.Logger.Fatal(e.Start(":" + port))
	e.Logger.Fatal(e.Start(":8080"))
}
