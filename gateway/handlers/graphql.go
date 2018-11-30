package handlers

import (
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"github.com/proelbtn/school-eve-navi/gateway/models"
	"net/http"
)

type Request struct {
	Query string `json:"query" query:"query"`
}

func GraphQLHandler(c echo.Context) (err error) {
	req := new(Request)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: models.GetQueryObject(),
	})

	if err != nil {
		c.Echo().Logger.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	if err = c.Bind(req); err != nil {
		return err
	}

	params := graphql.Params{
		Schema:        schema,
		RequestString: req.Query,
	}

	return c.JSON(http.StatusOK, graphql.Do(params))
}
