package models

import (
	"github.com/graphql-go/graphql"
)

type Souvenir struct {
	Name string
}

func GetSouvenirObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Souvenir",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(Souvenir); ok {
						return obj.Name, nil
					}
					return nil, nil
				},
			},
		},
	}

	return graphql.NewObject(config)
}
