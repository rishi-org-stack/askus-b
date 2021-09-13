package mgdb

import (
	user "askUs/v1/package/user"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserDB = "users"
const AuthDB = "auths"

type UserDb struct{}

// func (Utb *UserDb) GetUser(ctx context.Context, id string) (*user.UserAggregate, error) {
// 	var res = []user.UserAggregate{}
// 	db := ctx.Value("mgClient").(*mongo.Database)
// 	lookup := bson.D{{
// 		Key: "$lookup",
// 		Value: bson.D{{
// 			Key:   "from",
// 			Value: "auths",
// 		},
// 			{
// 				Key:   "localField",
// 				Value: "auth_id",
// 			},
// 			{
// 				Key:   "foreignField",
// 				Value: "_id",
// 			},
// 			{
// 				Key:   "as",
// 				Value: "auth",
// 			}},
// 	}}
// 	// unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$auth"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
// 	userCollection := db.Collection(UserDB)
// 	Id,err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return res, err
// 	}
// 	unwindStage := bson.D{{Key: "$eq", Value: bson.D{{Key: "_id", Value: Id}}}}

// 	cursor, err := userCollection.Aggregate(ctx, mongo.Pipeline{lookup, unwindStage})
// 	// tlookup:=bson.D{{
// 	// 	Key: "$match",
// 	// 	Value:bson.D{{
// 	// 		Key: "_id",
// 	// 	}}
// 	// }}
// 	// cursor, err := userCollection.Aggregate(ctx, mongo.Pipeline{lookup, unwindStage})
// 	if err != nil {
// 		return res, err
// 	}
// 	// var tt []bson.M
// 	if err = cursor.All(ctx, &res); err != nil {
// 		return res, err
// 	}
// 	return res, err
// }
func (udb UserDb) GetUser(ctx context.Context, id primitive.ObjectID) (*user.User, error) {
	db := ctx.Value("surround").(map[string]interface{})["mgClient"].(*mongo.Database)
	userColl := db.Collection(UserDB)
	us := &user.User{}
	result := userColl.FindOne(ctx, bson.D{{Key: "auth_id", Value: id}})
	err := result.Decode(us)
	if err != nil {
		return us, err
	}

	return us, nil
}
func (Utb UserDb) UpdateUser(ctx context.Context, us *user.User) (*user.User, error) {
	db := ctx.Value("surround").(map[string]interface{})["mgClient"].(*mongo.Database)

	_, err := db.Collection(UserDB).
		ReplaceOne(ctx, bson.D{{Key: "auth_id", Value: us.AuthID}},
			*us)

	return us, err
}
