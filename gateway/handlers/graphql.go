package handlers

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"github.com/proelbtn/school-eve-navi/gateway/models/queries"
)

type Request struct {
	Query          string                 `json:"query" query:"query"`
	OperationName  string                 `json:"operationName" query:"operationName"`
	VariableValues map[string]interface{} `json:"variables" query:"variables"`
}

func GraphQLHandler(c echo.Context) (err error) {
	req := new(Request)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queries.GetQueryObject(),
	})

	if err != nil {
		c.Echo().Logger.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	if err = c.Bind(req); err != nil {
		return err
	}

	params := graphql.Params{
		Schema:         schema,
		RequestString:  req.Query,
		VariableValues: req.VariableValues,
		OperationName:  req.OperationName,
	}

	return c.JSON(http.StatusOK, graphql.Do(params))
}
