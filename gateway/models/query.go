package models

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/proelbtn/school-eve-navi/gateway/protos/food"
	"google.golang.org/grpc"
)

func GetFoodObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Food",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*food.FoodResponse); ok {
						return obj.Name, nil
					}
					return nil, errors.New("unable to cast")
				},
			},
		},
	}

	return graphql.NewObject(config)
}

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

func GetFoodField() *graphql.Field {
	return &graphql.Field{
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
	}
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
		Type: GetShopObject(),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return nil, nil
		},
	}
}

func GetQueryObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"foods":     GetFoodField(),
			"shops":     GetShopField(),
			"souvenirs": GetSouvenirField(),
		},
	}

	return graphql.NewObject(config)
}
