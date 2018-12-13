package resolvers

import (
	"github.com/proelbtn/school-eve-navi/gateway/protos/food"
	"google.golang.org/grpc"
)

func GetFoodClient() (food.FoodClient, error) {
	cc, err := grpc.Dial("food:80", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := food.NewFoodClient(cc)

	return client, nil
}
