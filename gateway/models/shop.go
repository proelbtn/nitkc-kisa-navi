package models

import (
	"github.com/graphql-go/graphql"
)

type Shop struct {
	Name string
}

func GetShopObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Shop",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(Shop); ok {
						return obj.Name, nil
					}
					return nil, nil
				},
			},
		},
	}

	return graphql.NewObject(config)
}
