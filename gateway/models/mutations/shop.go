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
			"price": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"genre_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			name, price, genreID := params.Args["name"], params.Args["price"], params.Args["genre_id"]

			client, err := resolvers.GetShopClient()
			if err != nil {
				return nil, err
			}

			req := &shop.ShopCreateQuery{
				Data: &shop.ShopData{
					Name:    name.(string),
					Price:   uint64(price.(int)),
					GenreId: uint64(genreID.(int)),
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
