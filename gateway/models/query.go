package models

import (
	"github.com/graphql-go/graphql"
)

func GetShopObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Shop",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
		},
	}

	return graphql.NewObject(config)
}

func GetSouvenirObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Souvenir",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return nil, nil
				},
			},
		},
	}

	return graphql.NewObject(config)
}

func GetShopField() *graphql.Field {
	return &graphql.Field{
		Type: GetShopObject(),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return nil, nil
		},
	}
}

func GetSouvenirField() *graphql.Field {
	return &graphql.Field{
		Type: GetSouvenirObject(),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return nil, nil
		},
	}
}

func GetQueryObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"foods":     GetFoodCategoryField(),
			"shops":     GetShopField(),
			"souvenirs": GetSouvenirField(),
		},
	}

	return graphql.NewObject(config)
}
