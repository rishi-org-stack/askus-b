package user

import (
	au "askUs/v1/package/auth"
	"askUs/v1/package/idea"
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	DB interface {
		GetUser(ctx context.Context, id primitive.ObjectID) (*User, error)
		UpdateUser(ctx context.Context, user *User) (*User, error)
	}
	Service interface {
		UpdateUser(ctx context.Context, user *User) (*response.Response, utilError.ApiErrorInterface)
		GetUser(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetIdea(ctx context.Context, id, pass string) (*response.Response, utilError.ApiErrorInterface)
		AddIdeaByID(ctx context.Context, idea *idea.Idea) (*response.Response, utilError.ApiErrorInterface)
		UpdateStatus(ctx context.Context, id, mark string) (*response.Response, utilError.ApiErrorInterface)
		UpdateIdea(ctx context.Context, ideainput *idea.Idea, id, pass string) (*response.Response, utilError.ApiErrorInterface)
		ExtendDeadline(ctx context.Context, id, yr, mon, hour string) (*response.Response, utilError.ApiErrorInterface)
	}
	//TODO:User ID needs to be of type Object Id
	User struct {
		ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name   string             `json:"name" bson:"name"`
		AuthID primitive.ObjectID `json:"authID" bson:"auth_id"`
		Ideas  map[string]*Status `json:"ideas" bson:"ideas"`
	}
	UserAggregate struct {
		User
		Auth au.AuthRequest `json:"auth"`
	}
	Idea   idea.Idea
	Status struct {
		MarkedAs  string `bson:"markedAs"`
		Deadline  string `bson:"deadline"`
		CreatedOn string `bson:"createdOn"`
	}
)

const (
	Ongoing   string = "ongoing"
	New       string = "New"
	Completed string = "Completed"
)
