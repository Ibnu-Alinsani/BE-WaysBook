package main

import (
	"fmt"
	// "os"
	"waysbook/database"
	"waysbook/pkg/postgresql"
	routes "waysbook/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{
			// "XMLHttpRequest",
			"X-Requested-With",
			"Content-Type",
			"Authorization",
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccept,
			echo.HeaderOrigin,
		},
	}))

	postgresql.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("api/v1"))
	// PORT := os.Getenv("PORT")

	fmt.Println("Server running in Localhost:5000")
	e.Logger.Fatal(e.Start(":4000"))
}
