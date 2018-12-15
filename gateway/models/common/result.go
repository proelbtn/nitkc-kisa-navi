package common

import (
	"sync"

	"github.com/graphql-go/graphql"
)

var object *graphql.Object
var once sync.Once

func GetResultObject() *graphql.Object {
	once.Do(func() {
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
	})

	return object
}
