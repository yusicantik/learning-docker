package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		fmt.Println("Ouch!")
		return c.HTML(http.StatusOK, "You're In")
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "77"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
