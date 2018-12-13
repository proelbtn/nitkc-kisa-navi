package queries

import (
	"github.com/graphql-go/graphql"
)

func GetQueryObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"foods":     GetFoodCategoryField(),
			"shops":     GetShopCategoryField(),
			"souvenirs": GetSouvenirCategoryField(),
		},
	}

	return graphql.NewObject(config)
}
