package user

import (
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"
	"github.com/lib/pq"
)

type (
	DB interface {
		FindOrCreateDoctor(ctx context.Context, email string) (*Doctor, error)
		FindOrCreatePatient(ctx context.Context, email string) (*Patient, error)
		CreateConnection(ctx context.Context, conn *Connections) (*Connections, error)
		GetDoctorByID(ctx context.Context, id float64) (*Doctor, error)
		GetDoctorByName(ctx context.Context, name string) (*Doctor, error)
		GetPatientByID(ctx context.Context, id float64) (*Patient, error)
		UpdatePatientByID(ctx context.Context, pt *Patient, id float64) (*Patient, error)
		UpdateDoctorByID(ctx context.Context, doc *Doctor, id float64) (*Doctor, error)
	}
	Service interface {
		FindOrCreateDoctor(ctx context.Context, email string) (*Doctor, utilError.ApiErrorInterface)
		FindOrCreatePatient(ctx context.Context, email string) (*Patient, utilError.ApiErrorInterface)
		GetDoctorByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetDoctorByName(ctx context.Context, name string) (*response.Response, utilError.ApiErrorInterface)
		GetPatientByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetUserByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		UpdatePatientByID(ctx context.Context, pt *Patient) (*response.Response, utilError.ApiErrorInterface)
		UpdateDoctortByID(ctx context.Context, doc *Doctor) (*response.Response, utilError.ApiErrorInterface)
	}
	//TODO:User ID needs to be of type Object Id
	Doctor struct {
		ID int `json:"id" gorm:"primary"`
		Info
		// Address
		Specialities  pq.StringArray `gorm:"type:text[]"`
		ExpInYears    string
		Experiences   *[]Experience
		DegreeID      int
		Degree        pq.StringArray `gorm:"type:text[]"`
		ConnectionsID int            `json:"connectionsID"`
		Connections   *Connections
	}
	Experience struct {
		ID            int `json:"id" gorm:"primary"`
		Institution   *Institution
		WorkedBetween string
		Department    string
		Title         string
		DoctorID      int `gorm:"not null"`
		Doctor        *Doctor
	}
	Institution struct {
		ID int `json:"id" gorm:"primary"`
		// Address
		Name         string
		ExperienceID int
		Experience   *Experience
	}
	Address struct {
		// ID      int `json:"id" gorm:"primary"`
		City    string
		Country string
		PinCode string
		State   string
		Street  string
	}
	Patient struct {
		ID int `json:"id" gorm:"primary"`
		Info
		// Address   *Address
		Symptoms      pq.StringArray `gorm:"type:text[]"`
		ConnectionsID int            `json:"connectionsID"`
		Connections   *Connections
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
		// Address
	}
	Connections struct {
		ID            int       `json:"id" gorm:"primary"`
		Doctors       []Doctor  `json:"doctors"`
		Patients      []Patient `json:"patients"`
		ConnectedWith string    `json:"connectedWith"`
	}
)

const (
	DoctorClient  = "doctor"
	PatientClient = "patient"
)

const (
	Onboarded = "ONBOARDED"
	NEW       = "NEW"
	Dropped   = "DROPPED"
)

const Schema = "usr"

func (Doctor) TableName() string {
	return Schema + ".doctors"
}

func (Patient) TableName() string {
	return Schema + ".patients"
}

func (Institution) TableName() string {
	return Schema + ".institutions"
}

func (Experience) TableName() string {
	return Schema + ".experiences"
}

func (Connections) TableName() string {
	return Schema + ".connections"
}
