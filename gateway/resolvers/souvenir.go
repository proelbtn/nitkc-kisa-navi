package resolvers

import (
	"github.com/proelbtn/school-eve-navi/gateway/protos/souvenir"
	"google.golang.org/grpc"
)

func GetSouvenirClient() (souvenir.SouvenirClient, error) {
	cc, err := grpc.Dial("souvenir:80", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := souvenir.NewSouvenirClient(cc)

	return client, nil
}
