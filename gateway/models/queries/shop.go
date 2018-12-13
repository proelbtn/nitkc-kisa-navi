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
				Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
					client, err := resolvers.GetShopClient()
					if err != nil {
						return nil, err
					}

					res, err := client.Search(context.Background(), &shop.ShopSearchQuery{
						Name:      "test",
						Genre:     3,
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

func GetShopCategoryField() *graphql.Field {
	return &graphql.Field{
		Type: GetShopCategoryObject(),
		Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
			return true, nil
		},
	}
}
