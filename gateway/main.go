package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/proelbtn/school-eve-navi/gateway/handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	// ref: https://graphql.org/graphql-js/graphql-clients/
	allowedMethods := []string{http.MethodGet, http.MethodPost}
	e.Match(allowedMethods, "/graphql", handlers.GraphQLHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
