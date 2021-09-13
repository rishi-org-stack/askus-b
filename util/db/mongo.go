package db

// import (
// 	"context"
// 	"fmt"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"logit/v1/util/config"
// )

// func Connect(ctx context.Context, env *config.Env) *mongo.Database {
// 	client, err := mongo.Connect(
// 		ctx,
// 		options.Client().
// 			ApplyURI(env.DB),
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Successfully connected and pinged.")
// 	// defer func() {
// 	// 	if err = client.Disconnect(ctx); err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }()
// 	//TODO: ADD DB NAME FROM ENV TYPE
// 	return (client.Database("LogIt"))
// }
