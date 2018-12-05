package models

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/proelbtn/school-eve-navi/gateway/protos/food"
)

func GetFoodGenreObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "FoodGenre",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*Genre); ok {
						return obj.id, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*Genre); ok {
						return obj.name, nil
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
					arr := make([]*Genre, 0)

					for k, v := range food.FoodEntity_Genre_name {
						arr = append(arr, &Genre{id: int64(k), name: v})
					}

					return arr, nil
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
