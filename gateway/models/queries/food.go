package queries

import (
	"context"
	"errors"

	"github.com/proelbtn/school-eve-navi/gateway/resolvers"

	"github.com/graphql-go/graphql"
	"github.com/proelbtn/school-eve-navi/gateway/protos/food"
)

func GetFoodRecordObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "FoodRecord",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*food.FoodRecord); ok {
						return obj.Id, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*food.FoodRecord); ok {
						return obj.Data.Name, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
		},
	}

	return graphql.NewObject(config)
}

func GetFoodGenreObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "FoodGenre",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*food.FoodGenre); ok {
						return obj.Id, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*food.FoodGenre); ok {
						return obj.Name, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
		},
	}

	return graphql.NewObject(config)
}

func GetFoodCategoryObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "FoodCategory",
		Fields: graphql.Fields{
			"genres": &graphql.Field{
				Type: graphql.NewList(GetFoodGenreObject()),
				Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
					client, err := resolvers.GetFoodClient()
					if err != nil {
						return nil, err
					}

					res, err := client.GetGenres(context.Background(), &food.FoodEmptyQuery{})
					if err != nil {
						return nil, err
					}

					return res.Genres, nil
				},
			},
			"records": &graphql.Field{
				Type: graphql.NewList(GetFoodRecordObject()),
				Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
					client, err := resolvers.GetFoodClient()
					if err != nil {
						return nil, err
					}

					res, err := client.Search(context.Background(), &food.FoodSearchQuery{
						Name:      "test",
						Genre:     1,
						Latitude:  0.123,
						Longitude: 0.123,
					})
					if err != nil {
						return nil, err
					}

					return res.Records, nil
				},
			},
		},
	}

	return graphql.NewObject(config)
}

func GetFoodCategoryField() *graphql.Field {
	return &graphql.Field{
		Type: GetFoodCategoryObject(),
		Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
			return true, nil
		},
	}
}
