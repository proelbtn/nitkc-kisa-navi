package queries

import (
	"context"
	"errors"

	"github.com/proelbtn/school-eve-navi/gateway/resolvers"

	"github.com/graphql-go/graphql"
	"github.com/proelbtn/school-eve-navi/gateway/protos/shop"
)

func GetShopRecordObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "ShopRecord",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*shop.ShopRecord); ok {
						return obj.Id, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*shop.ShopRecord); ok {
						return obj.Data.Name, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
		},
	}

	return graphql.NewObject(config)
}

func GetShopGenreObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "ShopGenre",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*shop.ShopGenre); ok {
						return obj.Id, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*shop.ShopGenre); ok {
						return obj.Name, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
		},
	}

	return graphql.NewObject(config)
}

func GetShopCategoryObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "ShopCategory",
		Fields: graphql.Fields{
			"genres": &graphql.Field{
				Type: graphql.NewList(GetShopGenreObject()),
				Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
					client, err := resolvers.GetShopClient()
					if err != nil {
						return nil, err
					}

					res, err := client.GetGenres(context.Background(), &shop.ShopEmptyQuery{})
					if err != nil {
						return nil, err
					}

					return res.Genres, nil
				},
			},
			"records": &graphql.Field{
				Type: graphql.NewList(GetShopRecordObject()),
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
					"genre_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, long, lat, name, genreID := params.Args["id"], params.Args["longitude"], params.Args["latitude"], params.Args["name"], params.Args["genre_id"]

					client, err := resolvers.GetShopClient()
					if err != nil {
						return nil, err
					}

					if id != nil && long == nil && lat == nil && name == nil && genreID == nil {
						res, err := client.Get(context.Background(), &shop.ShopGetQuery{
							Id: uint64(id.(int)),
						})
						if err != nil {
							return nil, err
						}

						return []*shop.ShopRecord{res.Record}, nil
					} else if id == nil {
						query := shop.ShopSearchQuery{}

						if name != nil {
							query.Name = name.(string)
						}
						if genreID != nil {
							query.GenreId = uint64(genreID.(int))
						}

						res, err := client.Search(context.Background(), &query)
						if err != nil {
							return nil, err
						}

						// TODO: fitlering

						return res.Records, nil
					}

					return nil, errors.New("unexpected call")
				},
			},
		},
	}

	return graphql.NewObject(config)
}

func GetShopCategoryField() *graphql.Field {
	return &graphql.Field{
		Type: GetShopCategoryObject(),
		Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
			return true, nil
		},
	}
}
