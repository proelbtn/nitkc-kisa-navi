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
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"genre": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Type: graphql.NewList(GetFoodRecordObject()),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, long, lat, name, genre := params.Args["id"], params.Args["longitude"], params.Args["latitude"], params.Args["name"], params.Args["genre"]

					client, err := resolvers.GetFoodClient()
					if err != nil {
						return nil, err
					}

					if id != nil && long == nil && lat == nil && name == nil && genre == nil {
						res, err := client.Get(context.Background(), &food.FoodGetQuery{
							Id: uint64(id.(int)),
						})
						if err != nil {
							return nil, err
						}

						return []*food.FoodRecord{res.Record}, nil
					} else if id == nil && long != nil && lat != nil {
						query := food.FoodSearchQuery{
							Longitude: long.(float64),
							Latitude:  lat.(float64),
						}

						if name != nil {
							query.Name = name.(string)
						}
						if genre != nil {
							query.Genre = uint64(genre.(int))
						}

						res, err := client.Search(context.Background(), &query)
						if err != nil {
							return nil, err
						}

						return res.Records, nil
					}

					return nil, errors.New("unexpected call")
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
