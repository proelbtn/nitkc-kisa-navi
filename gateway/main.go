package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/proelbtn/school-eve-navi/gateway/handlers"
)

func main() {
	e := echo.New()

	origins := os.Getenv("ALLOW_ORIGINS")

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(origins, ","),
	}))

	// ref: https://graphql.org/graphql-js/graphql-clients/
	allowedMethods := []string{http.MethodGet, http.MethodPost}
	e.Match(allowedMethods, "/graphql", handlers.GraphQLHandler)

	e.Logger.Fatal(e.Start("0.0.0.0:9000"))
}
