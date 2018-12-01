package models

import (
	"github.com/graphql-go/graphql"
	"github.com/proelbtn/school-eve-navi/gateway/protos/food"
)

func GetFoodObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Food",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(food.Food); ok {
						return obj.Name, nil
					}
					return nil, nil
				},
			},
		},
	}

	return graphql.NewObject(config)
}
