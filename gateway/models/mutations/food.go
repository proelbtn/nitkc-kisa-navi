package mutations

import (
	"github.com/graphql-go/graphql"
)

func CreateFood() *graphql.Field {
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
			userName, _ := params.Args["userName"].(string)
			description, _ := params.Args["description"].(string)
			photoURL, _ := params.Args["photoURL"].(string)
			email, _ := params.Args["email"].(string)

				panic(err)
			}
			infrastructure.NewUserRepository().Store(newUser)
			return newUser, nil
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
			userName, _ := params.Args["userName"].(string)
			description, _ := params.Args["description"].(string)
			photoURL, _ := params.Args["photoURL"].(string)
			email, _ := params.Args["email"].(string)

			newUser, err := user.NewUser(userName, description, photoURL, email)
			if err != nil {
				panic(err)
			}
			infrastructure.NewUserRepository().Store(newUser)
			return newUser, nil
		},
	}
}
