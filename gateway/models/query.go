package models

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/proelbtn/school-eve-navi/gateway/protos/food"
	"google.golang.org/grpc"
)

func GetQueryObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"foods": &graphql.Field{
				Type: GetFoodObject(),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					cc, err := grpc.Dial("shop:30001", grpc.WithInsecure())
					if err != nil {
						return nil, err
					}

					client := food.NewFoodClient(cc)
					req := food.FoodRequest{Name: "Hello, world!"}
					return client.Search(context.Background(), &req)
				},
			},
			"shops": &graphql.Field{
				Type: GetShopObject(),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return Shop{Name: "Shop"}, nil
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
