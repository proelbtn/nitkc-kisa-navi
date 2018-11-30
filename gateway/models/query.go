package models

import (
	"github.com/graphql-go/graphql"
)

func GetQueryObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"foods": &graphql.Field{
				Type: GetFoodObject(),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return Food{Name: "Food"}, nil
				},
			},
			"shops": &graphql.Field{
				Type: GetShopObject(),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return Food{Name: "Shop"}, nil
				},
			},
			"souvenirs": &graphql.Field{
				Type: GetSouvenirObject(),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return Souvenir{Name: "Souvenir"}, nil
				},
			},
		},
	}

	return graphql.NewObject(config)
}
