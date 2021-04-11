package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	var dbInstance = getInstance()
	e.GET("/", func(c echo.Context) error {
		var returned = dbInstance.getAll()
		fmt.Println(returned)
		return c.JSON(http.StatusOK, returned)
	})
	e.Logger.Fatal(e.Start(":1323"))
	dbInstance.closeConnection()
}