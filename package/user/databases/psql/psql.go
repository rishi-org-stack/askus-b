package psql

import (
	"askUs/v1/package/user"
	"context"
)

type UserDb struct{}

func (udb UserDb) GetUser(c context.Context, id int) (*user.User, error)

func (udb UserDb) UpdateUser(ctx context.Context, user *user.User) (*user.User, error)
