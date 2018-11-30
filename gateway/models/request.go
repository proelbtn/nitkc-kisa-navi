package models

type GraphQLRequest struct {
	Query string `json:"query" query:"query"`
}