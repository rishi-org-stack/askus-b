package auth

import (
	apiError "askUs/v1/util/error"
	apiRes "askUs/v1/util/response"
	"context"
)

//TODO: though the otp is gnratedfor user and vic versa it allowsgenration oftoken
type (
	DB interface {
		FindOrInsert(ctx context.Context, atr *AuthRequest) (*AuthRequest, error)
		Update(ctx context.Context, atr *AuthRequest) (*AuthRequest, error)
		// InsertUser(ctx context.Context, atr *AuthRequest) (AuthRequest{}, error)
		GetRequest(ctx context.Context, id int) (*AuthRequest, error)
	}
	Service interface {
		HandleAuth(ctx context.Context, atr *AuthRequest) (*apiRes.Response, apiError.ApiErrorInterface)
		Verify(ctx context.Context, otpReq *VerifyRequest) (*apiRes.Response, apiError.ApiErrorInterface)
		GetRequestByID(ctx context.Context, id int) (*AuthRequest, error) //apiError.ApiErrorInterface)
	}
	AuthRequest struct {
		ID     int    `json:"id"  gorm:"primaryKey"`
		Email  string `json:"email" gorm:"not null"`
		Status string `json:"status" gorm:"default:NEW"`
	}
	VerifyRequest struct {
		ID  int    `json:"id"`
		OTP string `json:"otp"`
	}
	StatusType string

	TokenGenratorInterface interface {
		GenrateToken(int, string, string) (string, error)
	}
	//DTO's
	AuthResponse struct {
		ID  int    `json:"id"`
		OTP string `json:"otp"`
	}
	// AuthRequest struct {
	// 	ID       primitive.ObjectID `json:"id,omitempty"`
	// 	Email    string             `json:"email,omitempty"`
	// 	Password string             `json:"password"`
	// }
)

const (
	New      StatusType = "New"
	Verified StatusType = "Verified"
	Invalid  StatusType = "Invalid"
	Old      StatusType = "Old"
)
const (
	DoctorClient  = "doctor"
	PatientClient = "patient"
)

func (AuthRequest) TableName() string {
	return "auth.auth_requests"
}
