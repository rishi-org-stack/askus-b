package idea

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//I dont think for now we need route layer
type (
	Service interface {
		GetIdea(ctx context.Context, id, pass string) (*Idea, error)
		InsertIdea(ctx context.Context, idea Idea) (string, error)
		UpdateIdea(ctx context.Context, us *Idea) (*Idea, error)
	}
	DB interface {
		GetIdea(ctx context.Context, id primitive.ObjectID) (*Idea, error)
		InsertIdea(ctx context.Context, idea Idea) (primitive.ObjectID, error)
		UpdateIdea(ctx context.Context, us *Idea) (*Idea, error)
	}
	Idea struct {
		ID           primitive.ObjectID `bson:"_id,omitempty"`
		Name         string             `bson:"name,omitempty"`
		Description  string             `bson:"description"`
		Type         string             `bson:"type"`
		ConcernedFor string             `bson:"concernedFor" json:"concernedFor"`
		Problem      string             `bson:"problem"`
		AccessKey    string             `bson:"accessKey,omitempty" json:"accessKey"`

		//Need to think about step
		//Steps
	}
)
