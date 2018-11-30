package main

import (
	"github.com/labstack/echo"
	"github.com/proelbtn/school-eve-navi/gateway/handlers"
	"net/http"
)

func main() {
	e := echo.New()

	// ref: https://graphql.org/graphql-js/graphql-clients/
	allowedMethods := []string{http.MethodGet, http.MethodPost}
	e.Match(allowedMethods, "/graphql", handlers.GraphQLHandler)

	e.Logger.Fatal(e.Start(":8081"))
}
