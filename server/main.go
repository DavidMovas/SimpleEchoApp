package main

import (
	"echoapp/db"
	"echoapp/handlers"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, err := db.NewDB()

	if err != nil {
		log.Fatal(err)
	}

	e.Static("", "client")

	registerHandlers(e, db)

	defer db.Close()

	err = e.Start(":8080")
	log.Fatal(err)
}

func registerHandlers(e *echo.Echo, db *db.DB) {
	handlers.NewUserHandler(e.Group("/users"), db)
}
