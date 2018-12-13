package mutations

import (
	"github.com/graphql-go/graphql"
)

func CreateSouvenir() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{ // 引数の定義
			"userName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"photoURL": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return false, nil
		},
	}
}

func DeleteSouvenir() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"userName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"photoURL": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return false, nil
		},
	}
}
