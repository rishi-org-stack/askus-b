package mgdb

import (
	"context"
	"askUs/v1/package/idea"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DB = "ideas"

type IdeaDB struct{}

func (udb IdeaDB) GetIdea(ctx context.Context, id primitive.ObjectID) (*idea.Idea, error) {
	db := ctx.Value("surround").(map[string]interface{})["mgClient"].(*mongo.Database)
	ideaColl := db.Collection(DB)
	ide := &idea.Idea{}
	result := ideaColl.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	err := result.Decode(ide)
	if err != nil {
		return ide, err
	}
	return ide, nil
}
func (udb IdeaDB) InsertIdea(ctx context.Context, idea idea.Idea) (primitive.ObjectID, error) {
	db := ctx.Value("surround").(map[string]interface{})["mgClient"].(*mongo.Database)
	ideaColl := db.Collection(DB)
	result, err := ideaColl.InsertOne(ctx, idea)
	if err != nil {
		return result.InsertedID.(primitive.ObjectID), err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (Utb IdeaDB) UpdateIdea(ctx context.Context, us *idea.Idea) (*idea.Idea, error) {
	db := ctx.Value("surround").(map[string]interface{})["mgClient"].(*mongo.Database)
	_, err := db.Collection(DB).
		ReplaceOne(ctx, bson.D{{Key: "_id", Value: us.ID}},
			us)

	return us, err
}
