package mgdb

// import (
// 	"context"
// 	"askUs/v1/package/auth"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// const DB = "auths"
// const UserDB = "users"

// type AuthDb struct{}

// func (auth AuthDb) ConnectToCollection() {

// }

// func (au AuthDb) FindOrInsert(ctx context.Context, atr *auth.AuthRequest) (interface{}, error) {
// 	var authReq = &auth.AuthRequest{}
// 	db := ctx.Value("mgClient").(*mongo.Database)
// 	err := db.Collection(DB).
// 		FindOne(ctx, bson.D{{
// 			"email", atr.Email,
// 		}}).
// 		Decode(authReq)
// 	if err != nil && err.Error() == "mongo: no documents in result" {
// 		res, err := db.Collection(DB).
// 			InsertOne(ctx, atr)
// 		return res.InsertedID, err
// 	}
// 	return authReq, nil
// }
// func (au AuthDb) InsertUser(ctx context.Context, atr *auth.AuthRequest) (interface{}, error) {
// 	db := ctx.Value("mgClient").(*mongo.Database)
// 	res, err := db.Collection(UserDB).InsertOne(ctx, bson.D{{Key: "auth_id", Value: atr.ID}})
// 	return res.InsertedID, err
// }
// func (au AuthDb) Update(ctx context.Context, atr *auth.AuthRequest) (interface{}, error) {
// 	db := ctx.Value("mgClient").(*mongo.Database)
// 	res, err := db.Collection(DB).
// 		ReplaceOne(ctx, bson.M{"_id": atr.ID}, atr)

// 	return res.UpsertedID, err
// }
// func (au AuthDb) GetRequest(ctx context.Context, id primitive.ObjectID) (*auth.AuthRequest, error) {
// 	db := ctx.Value("surround").(map[string]interface{})["mgClient"].(*mongo.Database)
// 	req := &auth.AuthRequest{}
// 	authColl := db.Collection(DB)
// 	result := authColl.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
// 	err := result.Decode(req)
// 	if err != nil {
// 		return &auth.AuthRequest{}, err
// 	}
// 	return req, nil
// }
