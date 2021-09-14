package user

import (
	"askUs/v1/package/auth"
	"askUs/v1/util"
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"
	"net/http"
)

const (
	source                 = "USER"
	source_auth            = "AUTH"
	source_idea            = "IDEA"
	ID_DECODE_ERROR        = source + "_ERROR_GEETING_ID"
	USER_GET_ERROR         = source + "_GET_ERROR"
	USER_GET_IDEA_ERROR    = source + "_" + source_idea + "_GET_ERROR"
	USER_AUTH_GET_ERROR    = source + "_" + source_auth + "_GET_ERROR"
	USER_COPY_ERROR        = source + "_COPY_ERROR"
	USER_IDEA_INSERT_ERROR = source + "_" + source_idea + "_INSERT_ERROR"
	USER_IDEA_UPDATE_ERROR = source + "_" + source_idea + "_INSERT_ERROR"
)

type (
	UserService struct {
		UserData    DB
		AuthService auth.Service
	}
)

func Init(db DB, authser auth.Service) Service {
	return &UserService{
		UserData:    db,
		AuthService: authser,
	}
}

func (uSer UserService) GetUser(ctx context.Context) (
	*response.Response,
	utilError.ApiErrorInterface) {

	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    ID_DECODE_ERROR,
		}
	}
	userRes, err := uSer.UserData.GetUser(ctx, Id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    USER_GET_ERROR,
		}
	}
	authID := userRes.AuthID
	authreq, err := uSer.AuthService.GetRequestByID(ctx, authID.Hex())
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    USER_AUTH_GET_ERROR,
		}
	}
	res := &UserAggregate{}
	res.Auth = *authreq
	res.User = *userRes
	return &response.Response{
		Status:  http.StatusOK,
		Message: "User Get success",
		Data:    res,
	}, nil
}

func (uSer UserService) UpdateUser(ctx context.Context, user *User) (*response.Response, utilError.ApiErrorInterface) {

	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    ID_DECODE_ERROR,
		}
	}
	userdbStruct, err := uSer.UserData.GetUser(ctx, Id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    USER_GET_ERROR,
		}
	}
	util.TransferData(user, userdbStruct)
	_, err = uSer.UserData.UpdateUser(ctx, userdbStruct)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    USER_COPY_ERROR,
		}
	}

	return &response.Response{
		Status:  http.StatusFound,
		Message: "USer Update Succssfull",
		Data:    userdbStruct,
	}, nil
}
