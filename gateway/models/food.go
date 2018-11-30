package models

import (
	"github.com/graphql-go/graphql"
)

type Food struct {
	Name string
}

func GetFoodObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Food",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(Food); ok {
						return obj.Name, nil
					}
					return nil, nil
				},
			},
		},
	}

	return graphql.NewObject(config)
}
