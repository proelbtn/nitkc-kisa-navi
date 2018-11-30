package handlers

import (
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"net/http"
)

type Shop struct {
	Name string
}

func getSchema() (graphql.Schema, error) {
	shop := graphql.NewObject(graphql.ObjectConfig{
		Name: "Shop",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if shop, ok := p.Source.(Shop); ok {
						return shop.Name, nil
					}
					return p.Source, nil
				},
			},
		},
	})

	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"shops": &graphql.Field{
				Type: shop,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return Shop{Name: "Hello, World"}, nil
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: query,
	})
}

type Request struct {
	Query string `json:"query" query:"query"`
}

func GraphQLHandler(c echo.Context) (err error) {
	req := new(Request)
	schema, err := getSchema()

	if err != nil {
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
