package advice

import (
	"askUs/v1/package/user"
	"askUs/v1/util"
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"
	"fmt"
	"net/http"
)

const (
	source            = "ADVICE"
	USER_UNAUTHORIZED = source + "_USER_UNAUTHORZED"
	ADVICE_GET_ERROR  = source + "_GET_ERROR"
)

type AdviceService struct {
	adviceData  DB
	userservice User
}

func Init(db DB, user User) Service {
	return &AdviceService{
		adviceData:  db,
		userservice: user,
	}
}
func (adServ AdviceService) CreateAdvice(ctx context.Context, adv *Advice) (*response.Response, utilError.ApiErrorInterface) {
	id, _ := util.GetFromServiceCtx(ctx, "id").(float64)
	clientType := util.GetFromServiceCtx(ctx, "userType")
	fmt.Println(clientType)
	if clientType == DoctorClient {
		adv.PostedBy = int(id)
		adv.Type = GLOBAL
		adv.PostedFor = 0
		res, err := adServ.adviceData.CreateAdvice(ctx, adv)
		if err != nil {
			return &response.Response{}, utilError.ApiError{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Code:    "internal server error",
			}
		}
		return &response.Response{
			Data:    res,
			Message: "Create adice success",
			Status:  http.StatusCreated,
		}, nil
	}
	return &response.Response{}, utilError.ApiError{
		Status:  http.StatusUnauthorized,
		Message: "Patients are no allowed to create advice",
		Code:    USER_UNAUTHORIZED,
	}
}
func (adServ AdviceService) CreatePersonelAdvice(ctx context.Context, adv *Advice, pt string) (*response.Response, utilError.ApiErrorInterface) {
	id, _ := util.GetFromServiceCtx(ctx, "id").(float64)
	ptID, _ := util.StringToInt(pt)
	// fmt.Println(pt)
	fmt.Println(ptID)
	clientType := util.GetFromServiceCtx(ctx, "userType")
	if clientType == DoctorClient {
		adv.PostedBy = int(id)
		adv.Type = PERSONEL
		adv.PostedFor = ptID
		fmt.Println(adv)
		res, err := adServ.adviceData.CreateAdvice(ctx, adv)
		if err != nil {
			return &response.Response{}, utilError.ApiError{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Code:    "internal server error",
			}
		}
		return &response.Response{
			Data:    res,
			Message: "Create advice success",
			Status:  http.StatusCreated,
		}, nil
	}
	return &response.Response{}, utilError.ApiError{
		Status:  http.StatusUnauthorized,
		Message: "Patients are no allowed to create advice",
		Code:    USER_UNAUTHORIZED,
	}
}

// func (adServ AdviceService) GetGlobalAdvices(ctx context.Context) (*response.Response, utilError.ApiErrorInterface)
func (adServ AdviceService) GetGlobalAdvice(ctx context.Context, id string) (*response.Response, utilError.ApiErrorInterface) {
	Id, _ := util.StringToInt(id)
	adv, err := adServ.adviceData.GetGlobalAdviceByID(ctx, float64(int64(Id)))
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    ADVICE_GET_ERROR,
		}

	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Get Glbal advice success",
		Data:    adv,
	}, nil
}
func (adServ AdviceService) GetPersonelAdvices(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id, _ := util.GetFromServiceCtx(ctx, "id").(float64)
	adv, err := adServ.adviceData.GetAllPersonelAdvices(ctx, float64(int64(id)))
	if err != nil {
		return &response.Response{}, utilError.ApiError{

			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    ADVICE_GET_ERROR,
		}

	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Get Glbal advice success",
		Data:    adv,
	}, nil
}

func (adServ AdviceService) GetPersonelAdvicesPostedByMe(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id, _ := util.GetFromServiceCtx(ctx, "id").(float64)
	adv, err := adServ.adviceData.GetAllPersonelAdvicesPostedByDoc(ctx, float64(int64(id)))
	if err != nil {
		return &response.Response{}, utilError.ApiError{

			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    ADVICE_GET_ERROR,
		}

	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Get Glbal advice success",
		Data:    adv,
	}, nil
}
func (adServ AdviceService) GetPersonelAdvice(ctx context.Context, id string) (*response.Response, utilError.ApiErrorInterface) {
	Id, _ := util.StringToInt(id)
	patientID, _ := util.GetFromServiceCtx(ctx, "id").(float64)
	adv, err := adServ.adviceData.GetPersonelAdviceByID(ctx, float64(int64(Id)))
	if err != nil {
		return &response.Response{}, utilError.ApiError{

			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    ADVICE_GET_ERROR,
		}

	}
	if adv.PostedFor != int(patientID) {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusUnauthorized,
			Message: "your id doesn't matches with id of advice",
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Get Glbal advice success",
		Data:    adv,
	}, nil
}
func (adServ AdviceService) GetDocAdvices(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	docID, _ := util.GetFromServiceCtx(ctx, "id").(float64)
	adv, err := adServ.adviceData.GetAllDocAdvices(ctx, float64(int64(docID)))
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    ADVICE_GET_ERROR,
		}

	}

	return &response.Response{
		Status:  http.StatusOK,
		Message: "Get all  advices of a particular doc success",
		Data:    adv,
	}, nil
}
func (adServ AdviceService) LikeAdvice(ctx context.Context, advID string) (*response.Response, utilError.ApiErrorInterface) {
	ID, _ := util.GetFromServiceCtx(ctx, "id").(float64)
	advid, _ := util.StringToInt(advID)
	like := &Like{}
	like.AdviceID = advid
	like.LikedBy = int(ID)
	like, err := adServ.adviceData.CreateLike(ctx, like)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Get all  advices of a particular doc success",
		Data:    like,
	}, nil
}

func (adserv AdviceService) GetPatientAndMyAdvices(ctx context.Context, pid string) (*response.Response, utilError.ApiErrorInterface) {
	ID, _ := util.GetFromServiceCtx(ctx, "id").(float64)
	result, err := adserv.userservice.GetPatientByID(ctx, pid)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	var data = &PatientAdiceGRP{}
	data.Patient = *result.Data.(*user.Patient)
	advices, err := adserv.adviceData.GetAdviceWithPIDAndDID(ctx, ID, pid)
	if err != nil {
		return &response.Response{},
			utilError.ApiError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
	}
	data.Advices = *advices
	return &response.Response{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}
