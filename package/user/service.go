package user

import (
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"
)

type (
	DB interface {
		GetUser(ctx context.Context, id int) (*User, error)
		UpdateUser(ctx context.Context, user *User) (*User, error)
	}
	Service interface {
		UpdateUser(ctx context.Context, user *User) (*response.Response, utilError.ApiErrorInterface)
		GetUser(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
	}
	//TODO:User ID needs to be of type Object Id
	Doctor struct {
		ID int `json:"id" gorm:"primary"`
		Info
		Specialities      []string
		Experience        string
		DegreeID          int
		Degree            *[]Degree
		Patients          *[]Patient
		DoctorsConnection *[]Doctor
	}

	Patient struct {
		ID int `json:"id" gorm:"primary"`
		Info
		Symptoms  []string
		Following *[]Doctor
	}
	Degree struct {
	}
	Info struct {
		Email string
		Phone string
		Name  string `json:"name" gorm:"not null"`
		Age   int
		State string
		Sex   string
	}
)

const (
	Onboarded = "ONBOARDED"
	NEW       = "NEW"
	Dropped   = "DROPPED"
)
