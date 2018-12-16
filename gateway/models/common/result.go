package common

import (
	"github.com/graphql-go/graphql"
)

var object *graphql.Object

func GetResultObject() *graphql.Object {
	if object == nil {
		config := graphql.ObjectConfig{
			Name: "Result",
			Fields: graphql.Fields{
				"success": &graphql.Field{
					Type: graphql.Boolean,
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						return params.Source.(bool), nil
					},
				},
			},
		}

		object = graphql.NewObject(config)
	}

	return object
}
