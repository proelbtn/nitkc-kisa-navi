package mutations

import (
	"github.com/graphql-go/graphql"
)

func GetMutationObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createFood":     CreateFood(),
			"deleteFood":     DeleteFood(),
			"createShop":     CreateShop(),
			"deleteShop":     DeleteShop(),
			"createSouvenir": CreateSouvenir(),
			"deleteSouvenir": DeleteSouvenir(),
		},
	}

	return graphql.NewObject(config)
}
