package user

import (
	"askUs/v1/package/auth"
	"askUs/v1/package/idea"
	"askUs/v1/util"
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
		IdeaService idea.Service
		AuthService auth.Service
	}
)

func Init(db DB, authser auth.Service, ideaSer idea.Service) Service {
	return &UserService{
		UserData:    db,
		AuthService: authser,
		IdeaService: ideaSer,
	}
}

func (uSer UserService) GetUser(ctx context.Context) (*response.Response,
	utilError.ApiErrorInterface) {
	Id, err := util.StringToObjectID(ctx)
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
	Id, err := util.StringToObjectID(ctx)
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
func (uSer UserService) GetIdea(ctx context.Context, id, pass string) (*response.Response, utilError.ApiErrorInterface) {
	Id, err := util.StringToObjectID(ctx)
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
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    USER_GET_ERROR,
		}
	}
	if len(userdbStruct.Ideas) == 0 {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "No ideas available",
			Code:    USER_GET_IDEA_ERROR,
		}

	}
	if userdbStruct.Ideas[id] == nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "Ide with prticular id is not available",
			Code:    USER_GET_IDEA_ERROR,
		}
	}
	userIdea, err := uSer.IdeaService.GetIdea(ctx, id, pass)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    USER_GET_IDEA_ERROR,
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Get Idea SuccessFull",
		Data:    userIdea,
	}, nil
}

func (uSer UserService) AddIdeaByID(ctx context.Context, idea *idea.Idea) (*response.Response, utilError.ApiErrorInterface) {
	Id, err := util.StringToObjectID(ctx)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    ID_DECODE_ERROR,
		}
	}
	fmt.Println(idea)
	ideaID, err := uSer.IdeaService.InsertIdea(ctx, *idea)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    USER_IDEA_INSERT_ERROR,
		}
	}
	userdbStruct, err := uSer.UserData.GetUser(ctx, Id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    USER_GET_ERROR,
		}
	}
	if userdbStruct.Ideas == nil {
		userdbStruct.Ideas = make(map[string]*Status)
	}
	userdbStruct.Ideas[ideaID] = &Status{
		CreatedOn: time.Now().Format(time.RFC3339),
		Deadline:  time.Now().Format(time.RFC3339),
		MarkedAs:  "New",
	}
	_, err = uSer.UserData.UpdateUser(ctx, userdbStruct)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    ID_DECODE_ERROR,
		}
	}
	return &response.Response{
		Status:  http.StatusCreated,
		Message: "Create Idea Success",
		Data:    userdbStruct,
	}, nil
}

func (uSer UserService) UpdateIdea(ctx context.Context, ideainput *idea.Idea, id, pass string) (*response.Response, utilError.ApiErrorInterface) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    ID_DECODE_ERROR,
		}
	}
	ideainput.ID = Id
	userIdea, err := uSer.IdeaService.GetIdea(ctx, id, pass)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    USER_GET_IDEA_ERROR,
		}
	}
	err = util.TransferData(ideainput, userIdea)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    USER_GET_IDEA_ERROR,
		}
	}
	updatedIdea, err := uSer.IdeaService.UpdateIdea(ctx, userIdea)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    USER_IDEA_UPDATE_ERROR,
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Update Idea successfull",
		Data:    updatedIdea,
	}, nil
}
func (uSer UserService) UpdateStatus(ctx context.Context, id, mark string) (*response.Response, utilError.ApiErrorInterface) {
	uId, err := util.StringToObjectID(ctx)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    ID_DECODE_ERROR,
		}
	}
	user, err := uSer.UserData.GetUser(ctx, uId)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    USER_GET_ERROR,
		}
	}
	switch mark {
	case "ongoing":
		user.Ideas[id].MarkedAs = Ongoing
	case "completed":
		user.Ideas[id].MarkedAs = Completed
	}
	res, err := uSer.UpdateUser(ctx, user)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    USER_IDEA_UPDATE_ERROR,
		}
	}
	user = res.Data.(*User)
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Update Status successfull",
		Data:    user,
	}, nil
}
func (uSer UserService) ExtendDeadline(ctx context.Context, id, yr, mon, hour string) (*response.Response, utilError.ApiErrorInterface) {
	yri, err := strconv.Atoi(yr)
	moni, err := strconv.Atoi(mon)
	houri, err := strconv.Atoi(hour)
	uid := util.GetFromServiceCtx(ctx, "id").(string)
	uId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    ID_DECODE_ERROR,
		}
	}
	user, err := uSer.UserData.GetUser(ctx, uId)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Code:    USER_GET_ERROR,
		}
	}
	user.Ideas[id].Deadline = time.Now().
		AddDate(yri, moni, houri).
		Format(time.RFC3339)
	res, err := uSer.UpdateUser(ctx, user)
	user = res.Data.(*User)
	return &response.Response{}, utilError.ApiError{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
		Code:    USER_IDEA_UPDATE_ERROR,
	}
}
