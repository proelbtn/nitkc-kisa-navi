package handlers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/proelbtn/school-eve-navi/gateway/models"
)

func GraphQLHandler(c echo.Context) (err error) {
	req := new(models.GraphQLRequest)

	if err = c.Bind(req); err != nil {
		return err
	}

	return c.String(http.StatusOK, req.Query)
}