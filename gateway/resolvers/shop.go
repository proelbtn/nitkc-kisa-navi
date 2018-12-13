package resolvers

import (
	"github.com/proelbtn/school-eve-navi/gateway/protos/shop"
	"google.golang.org/grpc"
)

func GetShopClient() (shop.ShopClient, error) {
	cc, err := grpc.Dial("shop:80", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := shop.NewShopClient(cc)

	return client, nil
}
