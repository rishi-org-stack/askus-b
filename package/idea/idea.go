package idea

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IdeaService struct {
	IdeaData DB
}

func Init(db DB) Service {
	return &IdeaService{
		IdeaData: db,
	}
}

func (iSr IdeaService) GetIdea(ctx context.Context, id, pass string) (*Idea, error) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &Idea{}, err
	}

	ideaRes, err := iSr.IdeaData.GetIdea(ctx, Id)
	if err != nil {
		return ideaRes, err
	}
	if ideaRes.AccessKey != "" {
		allowed := iSr.authenticate(ideaRes, pass)
		if !allowed {
			return &Idea{}, fmt.Errorf("authenticaton fail for current idea can not be accessed")
		}
	}
	return ideaRes, nil
}

func (iSr IdeaService) InsertIdea(ctx context.Context, idea Idea) (string, error) {
	ideaRes, err := iSr.IdeaData.InsertIdea(ctx, idea)
	if err != nil {
		return ideaRes.Hex(), err
	}
	return ideaRes.Hex(), nil
}

func (iSr IdeaService) UpdateIdea(ctx context.Context, us *Idea) (*Idea, error) {
	ideaRes, err := iSr.IdeaData.UpdateIdea(ctx, us)
	if err != nil {
		return ideaRes, err
	}
	return ideaRes, nil
}

//TODO: improve
func (iSr IdeaService) authenticate(idea *Idea, pass string) bool {
	if idea.AccessKey == pass {
		return true
	}
	return false
}
