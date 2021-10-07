package advice

import (
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"

	"github.com/lib/pq"
)

type (
	DB interface {
		CreateAdvice(context.Context, *Advice) (*Advice, error)
		CreateLike(context.Context, *Like) (*Like, error)
		GetGlobalAdviceByID(context.Context, float64) (*Advice, error)
		GetPersonelAdviceByID(context.Context, float64) (*Advice, error)
		GetAdviceByPatientID(context.Context, float64) (*Advice, error)
		GetAdviceByDocID(context.Context, float64) (*Advice, error)
		GetAllPersonelAdvices(context.Context, float64) (*[]Advice, error)
		GetAllDocAdvices(context.Context, float64) (*[]Advice, error)
	}
	Service interface {
		CreateAdvice(ctx context.Context, adv *Advice) (*response.Response, utilError.ApiErrorInterface)
		CreatePersonelAdvice(ctx context.Context, adv *Advice, ptId string) (*response.Response, utilError.ApiErrorInterface)
		// GetGlobalAdvices(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetGlobalAdvice(ctx context.Context, id string) (*response.Response, utilError.ApiErrorInterface)
		GetPersonelAdvices(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		GetPersonelAdvice(ctx context.Context, id string) (*response.Response, utilError.ApiErrorInterface)
		GetDocAdvices(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
		LikeAdvice(ctx context.Context, advId string) (*response.Response, utilError.ApiErrorInterface)
	}
	Advice struct {
		ID             int            `gorm:"primary" json:"ID"`
		Heading        string         `json:"heading"`
		Body           string         `json:"body"`
		LikedByPatient []Like         `json:"likedByPatient"`
		LikedByDoc     []Like         `json:"likedByDoc"`
		Tags           pq.StringArray `json:"tags" gorm:"type: text[]"`
		PostedBy       int            `json:"postedBy"`
		Type           string         `json:"type"`
		PostedFor      int            `json:"postedFor"`
	}

	Like struct {
		ID       int `gorm:"primary" json:"ID"`
		LikedBy  int `json:"likedBy"`
		AdviceID int `json:"adviceID"`
		Advice   *Advice
	}
)

const (
	DoctorClient  = "doctor"
	PatientClient = "patient"
)

const (
	GLOBAL   = "GLOBAL"
	PERSONEL = "PERSONEL"
	schema   = "advice"
)

func (Advice) TableName() string {
	return schema + ".advices"
}
func (Like) TableName() string {
	return schema + ".likes"
}
