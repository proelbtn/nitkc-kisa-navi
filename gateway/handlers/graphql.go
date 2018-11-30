package handlers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/proelbtn/school-eve-navi/gateway/models"

	"encoding/json"
	"github.com/graphql-go/graphql"
)

func GraphQLHandler(c echo.Context) (err error) {
	fields := graphql.Fields {
		"hello": &graphql.Field {
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig {
		Name: "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig {
		Query: graphql.NewObject(rootQuery),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	req := new(models.GraphQLRequest)

	if err = c.Bind(req); err != nil {
		return err
	}

	params := graphql.Params {
		Schema: schema,
		RequestString: req.Query,
	}

	res, _ := json.Marshal(graphql.Do(params))

	return c.String(http.StatusOK, string(res))
}