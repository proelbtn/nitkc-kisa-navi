package mutations

import (
	"github.com/graphql-go/graphql"
)

func GetMutationObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createFoods":     CreateFoods(),
			"deleteFoods":     DeleteFoods(),
			"createShops":     CreateShops(),
			"deleteShops":     DeleteShops(),
			"createSouvenirs": CreateSouvenirs(),
			"deleteSouvenirs": DeleteSouvenirs(),
		},
	}

	return graphql.NewObject(config)
}
