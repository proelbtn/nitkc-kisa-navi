package models

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/proelbtn/school-eve-navi/gateway/protos/food"
	"google.golang.org/grpc"
)

func GetQueryObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"foods": &graphql.Field{
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "Hello, World!",
						Description:  "A name of food",
					},
				},
				Type: GetFoodObject(),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, ok := p.Args["name"].(string)
					if !ok {
						return nil, errors.New("name must be string")
					}

					cc, err := grpc.Dial("localhost:30001", grpc.WithInsecure())
					if err != nil {
						return nil, err
					}

					client := food.NewFoodClient(cc)
					req := food.FoodRequest{Name: name}

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
