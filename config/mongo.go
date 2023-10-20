package config

import (
	"context"

	"github.com/qiniu/qmgo"
)

func NewMongoConn(ctx context.Context) *qmgo.Client {
	mongoURI := Envs.MONGO_URI
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: mongoURI})
	if err != nil {
		panic(err)
	}
	return client
}
