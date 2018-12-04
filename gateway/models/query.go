package models

import (
	"context"
	"errors"

	"google.golang.org/grpc"

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
			// validate params
			name, ok := p.Args["name"].(string)
			if !ok {
				return nil, errors.New("name must be string")
			}

			// create request
			req := food.FoodRequest{Name: name}

			// prepare connection to upstream
			cc, err := grpc.Dial("food:80", grpc.WithInsecure())
			if err != nil {
				return nil, err
			}
			client := food.NewFoodClient(cc)

			// resolve
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
			"foods":     GetFoodField(),
			"shops":     GetShopField(),
			"souvenirs": GetSouvenirField(),
		},
	}

	return graphql.NewObject(config)
}
