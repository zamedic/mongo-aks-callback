package mongo_aks_oidc

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	aksCallback := NewAksCallback(WithScope("a05cf8a1-758d-4d14-8c56-7e89741b558f/.default"))
	uri := "mongodb+srv://xxxxxx.mongodb.net/?authSource=$external&authMechanism=MONGODB-OIDC"
	opts := options.Client().ApplyURI(uri)
	callback, err := aksCallback.GetAksCallback()
	if err != nil {
		panic(err)
	}
	opts.Auth.OIDCMachineCallback = callback
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.Disconnect(context.Background())
	}()

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

}
