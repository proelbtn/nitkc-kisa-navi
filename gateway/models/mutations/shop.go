package mutations

import (
	"context"

	"github.com/proelbtn/school-eve-navi/gateway/models/common"

	"github.com/graphql-go/graphql"
	"github.com/proelbtn/school-eve-navi/gateway/protos/shop"
	"github.com/proelbtn/school-eve-navi/gateway/resolvers"
)

func CreateShop() *graphql.Field {
	return &graphql.Field{
		Type: common.GetResultObject(),
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"genre_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"longitude": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"latitude": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"address": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"open": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"close": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			name, genreID, long, lat, address, open, close := params.Args["name"], params.Args["genre_id"], params.Args["longitude"], params.Args["latitude"], params.Args["address"], params.Args["open"], params.Args["close"]

			client, err := resolvers.GetShopClient()
			if err != nil {
				return nil, err
			}

			req := &shop.ShopCreateQuery{
				Data: &shop.ShopData{
					Name:      name.(string),
					GenreId:   uint64(genreID.(int)),
					Longitude: long.(float64),
					Latitude:  lat.(float64),
					Address:   address.(string),
					Open:      uint64(open.(int)),
					Close:     uint64(close.(int)),
				},
			}

			res, err := client.Create(context.Background(), req)
			if err != nil {
				return nil, err
			}

			return res.Success, nil
		},
	}
}

func DeleteShop() *graphql.Field {
	return &graphql.Field{
		Type: common.GetResultObject(),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"]

			client, err := resolvers.GetShopClient()
			if err != nil {
				return nil, err
			}

			req := &shop.ShopDeleteQuery{
				Id: uint64(id.(int)),
			}

			res, err := client.Delete(context.Background(), req)
			if err != nil {
				return nil, err
			}

			return res.Success, nil
		},
	}
}
