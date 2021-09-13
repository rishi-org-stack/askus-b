package auth

import (
	utilOTP "askUs/v1/util/auth"
	"askUs/v1/util/config"
	apiError "askUs/v1/util/error"
	apiRes "askUs/v1/util/response"
	"context"
	"net/http"
)

const (
	source                  = "AUTH"
	AUTH_INSERT_ERROR       = source + "_INSERT_ERROR"
	AUTH_BAD_REQUEST        = source + "_BAD_REQUEST"
	AUTH_OTP_INSERT_ERROR   = source + "_OTP_INSERT_ERROR"
	AUTH_UNAUTHORIZED_ERROR = source + "_INSERT_ERROR"
)

type AuthService struct {
	AuthData DB
	JwtSer   TokenGenratorInterface
	Config   *config.Env
}

var OTP string

func Init(db DB, js TokenGenratorInterface, config *config.Env) *AuthService {
	return &AuthService{
		AuthData: db,
		JwtSer:   js,
		Config:   config,
	}
}

func (authSer AuthService) HandleAuth(ctx context.Context) (*apiRes.Response, apiError.ApiErrorInterface) {
	atr := &AuthRequest{
		Email: "rishi@gmail.com",
	}
	res, err := authSer.AuthData.FindOrInsert(ctx, atr)
	if err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    AUTH_INSERT_ERROR,
		}
	}
	otp := utilOTP.GenrateOtp(authSer.Config.OTPExpiry)
	if err := otp.Set(res.ID); err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    AUTH_OTP_INSERT_ERROR,
		}
	}
	// OTP = otp.Otp
	return &apiRes.Response{
		Status:  http.StatusOK,
		Message: "Email authenticated",
		Data: &AuthResponse{
			ID:  res.ID,
			OTP: otp.Otp,
		},
	}, nil

}
func (authSer AuthService) Verify(ctx context.Context, otpReq *VerifyRequest) (*apiRes.Response, apiError.ApiErrorInterface) {
	otp := &utilOTP.OTP{}
	err := otp.Get(otpReq.ID)
	if err != nil {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    AUTH_BAD_REQUEST,
		}
	}
	if otp.Otp != otpReq.OTP {
		return &apiRes.Response{}, apiError.ApiError{
			Status:  http.StatusUnauthorized,
			Message: "otp doesn't matches",
			Code:    AUTH_UNAUTHORIZED_ERROR,
		}
	}
	return &apiRes.Response{
		Status:  http.StatusOK,
		Message: "otp  verified",
		Data:    "ok",
	}, nil
}
func (ar AuthService) GetRequestByID(ctx context.Context, id int) (*AuthRequest, error) {
	authR, err := ar.AuthData.GetRequest(ctx, 2)
	if err != nil {
		return &AuthRequest{}, nil
	}
	return authR, nil
}
