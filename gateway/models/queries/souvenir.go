package queries

import (
	"context"
	"errors"

	"github.com/proelbtn/school-eve-navi/gateway/resolvers"

	"github.com/graphql-go/graphql"
	"github.com/proelbtn/school-eve-navi/gateway/protos/souvenir"
)

func GetSouvenirRecordObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "SouvenirRecord",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*souvenir.SouvenirRecord); ok {
						return obj.Id, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*souvenir.SouvenirRecord); ok {
						return obj.Data.Name, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
		},
	}

	return graphql.NewObject(config)
}

func GetSouvenirGenreObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "SouvenirGenre",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*souvenir.SouvenirGenre); ok {
						return obj.Id, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if obj, ok := p.Source.(*souvenir.SouvenirGenre); ok {
						return obj.Name, nil
					}

					return nil, errors.New("unable to cast")
				},
			},
		},
	}

	return graphql.NewObject(config)
}

func GetSouvenirCategoryObject() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "SouvenirCategory",
		Fields: graphql.Fields{
			"genres": &graphql.Field{
				Type: graphql.NewList(GetSouvenirGenreObject()),
				Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
					client, err := resolvers.GetSouvenirClient()
					if err != nil {
						return nil, err
					}

					res, err := client.GetGenres(context.Background(), &souvenir.SouvenirEmptyQuery{})
					if err != nil {
						return nil, err
					}

					return res.Genres, nil
				},
			},
			"records": &graphql.Field{
				Type: graphql.NewList(GetSouvenirRecordObject()),
				Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
					client, err := resolvers.GetSouvenirClient()
					if err != nil {
						return nil, err
					}

					res, err := client.Search(context.Background(), &souvenir.SouvenirSearchQuery{
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

func GetSouvenirCategoryField() *graphql.Field {
	return &graphql.Field{
		Type: GetSouvenirCategoryObject(),
		Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
			return true, nil
		},
	}
}
