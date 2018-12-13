package mutations

import (
	"github.com/graphql-go/graphql"
)

func CreateFood() *graphql.Field {
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

func DeleteFood() *graphql.Field {
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
