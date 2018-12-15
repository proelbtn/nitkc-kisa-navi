package mutations

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/proelbtn/school-eve-navi/gateway/models/common"
	"github.com/proelbtn/school-eve-navi/gateway/protos/souvenir"
	"github.com/proelbtn/school-eve-navi/gateway/resolvers"
)

func CreateSouvenir() *graphql.Field {
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

			client, err := resolvers.GetSouvenirClient()
			if err != nil {
				return nil, err
			}

			req := &souvenir.SouvenirCreateQuery{
				Data: &souvenir.SouvenirData{
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

func DeleteSouvenir() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"]

			client, err := resolvers.GetSouvenirClient()
			if err != nil {
				return nil, err
			}

			req := &souvenir.SouvenirDeleteQuery{
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
