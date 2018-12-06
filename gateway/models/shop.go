package models

import (
	"context"
	"errors"

	"google.golang.org/grpc"

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

func GetShopCategoryObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "ShopCategory",
		Fields: graphql.Fields{
			"genres": &graphql.Field{
				Type: graphql.NewList(GetShopGenreObject()),
				Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
					arr := make([]*Genre, 0)

					for k, v := range shop.ShopGenre_name {
						if k == 0 {
							continue
						}

						arr = append(arr, &Genre{id: int64(k), name: v})
					}

					return arr, nil
				},
			},
			"records": &graphql.Field{
				Type: graphql.NewList(GetShopRecordObject()),
				Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
					cc, err := grpc.Dial("shop:80", grpc.WithInsecure())
					if err != nil {
						return nil, err
					}

					client := shop.NewShopClient(cc)

					res, err := client.Search(context.Background(), &shop.ShopSearchQuery{
						Name:      "test",
						Genre:     shop.ShopGenre_Invalid,
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
