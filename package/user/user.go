package user

import (
	"askUs/v1/package/report"
	"askUs/v1/util"
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"strconv"

	// "askUs/v1/util/response"
	"context"
	"net/http"

	"gorm.io/gorm"
)

const (
	source = "USER"

	ID_DECODE_ERROR           = source + "_ERROR_GEETING_ID"
	USER_GET_ERROR            = source + "_GET_ERROR"
	USER_DOCTOR_CREATE_ERROR  = source + "_DOCTOR_CREATE_ERROR"
	USER_PATIENT_CREATE_ERROR = source + "_PATIENT_CREATE_ERROR"
	USER_COPY_ERROR           = source + "_COPY_ERROR"
)

type (
	UserService struct {
		UserData DB
		// AuthService auth.Service
	}
)

func Init(db DB) (Service, report.User) {
	return &UserService{
			UserData: db,
			// AuthService: authser,
		},
		&UserService{
			UserData: db,
			// AuthService: authser,
		}
}

func (s UserService) FindOrCreateDoctor(ctx context.Context, email string) (*Doctor, utilError.ApiErrorInterface) {

	doc, err := s.UserData.FindOrCreateDoctor(ctx, email)

	if err != nil {
		return &Doctor{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "Error in creating or searching doctor",
			Code:    USER_DOCTOR_CREATE_ERROR,
		}
	}
	return doc, nil
}
func (s UserService) FindOrCreatePatient(ctx context.Context, email string) (*Patient, utilError.ApiErrorInterface) {
	doc, err := s.UserData.FindOrCreatePatient(ctx, email)
	if err != nil {
		return &Patient{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "Error in creating or searching doctor",
			Code:    USER_PATIENT_CREATE_ERROR,
		}
	}
	return doc, nil
}

func (s UserService) GetDoctorByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	doc, err := s.UserData.GetDoctorByID(ctx, id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    USER_GET_ERROR,
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Doctor retrieve successfull",
		Data:    doc,
	}, nil
}

func (s UserService) GetDoctorByName(ctx context.Context, name string) (*response.Response, utilError.ApiErrorInterface) {
	doc, err := s.UserData.GetDoctorByName(ctx, name)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "Error in retriving doctor from db",
			Code:    USER_GET_ERROR,
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Doctor retrieve successfull",
		Data:    doc,
	}, nil
}
func (s UserService) GetPatientByID(ctx context.Context, args ...string) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	if len(args) == 1 {
		if res, err := strconv.ParseFloat(args[0], 64); err == nil {
			id = res
		}
	}
	doc, err := s.UserData.GetPatientByID(ctx, id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "Error in retriving doctor from db",
			Code:    USER_GET_ERROR,
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "Doctor retrieve successfull",
		Data:    doc,
	}, nil
}
func (s UserService) GetUserByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	clientType := ctx.Value("surround").(map[string]interface{})["userType"].(string)
	if clientType == DoctorClient {
		return s.GetDoctorByID(ctx)
	}
	return s.GetPatientByID(ctx)
}

func (s UserService) UpdatePatientByID(ctx context.Context, pt *Patient) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	pt, err := s.UserData.UpdatePatientByID(ctx, pt, id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status: http.StatusBadRequest,
			// Code:    USER_PATIENT_UPDATE_ERROR,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "User Update success full",
		Data:    pt,
	}, nil
}

func (s UserService) UpdateDoctortByID(ctx context.Context, doc *Doctor) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	doc, err := s.UserData.UpdateDoctorByID(ctx, doc, id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status: http.StatusBadRequest,
			// Code:    USER_PATIENT_UPDATE_ERROR,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "User Update success full",
		Data:    doc,
	}, nil
}

//TODO: reques can e created by doctor to it self fix ths
func (s UserService) CreateReq(ctx context.Context, docid string) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	docID, _ := util.StringToInt(docid)
	req := &Request{
		SenderID:   int(id),
		DoctorID:   docID,
		Status:     PENDING,
		GenratedBy: ctx.Value("surround").(map[string]interface{})["userType"].(string),
	}
	_, err := s.UserData.GetReqWithSenderandDocID(ctx, id, docid, req.GenratedBy)
	if err != gorm.ErrRecordNotFound {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "already requesed",
		}
	}
	req, err = s.UserData.CreateReq(ctx, req)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadGateway,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status:  http.StatusCreated,
		Data:    req,
		Message: "user successfully requested to " + docid,
	}, nil
}
func (s UserService) GetMyRequests(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	reqs, err := s.UserData.GetReqsWithSenderID(ctx, id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadGateway,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status:  http.StatusCreated,
		Data:    reqs,
		Message: "get all requests  sent by current user successfull",
	}, nil
}

func (s UserService) GetRequestForMe(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	userType := util.GetFromServiceCtx(ctx, "userType")
	if userType == DoctorClient {
		reqs, err := s.UserData.GetReqsWithDocID(ctx, id)
		if err != nil {
			return &response.Response{}, utilError.ApiError{
				Status:  http.StatusBadGateway,
				Message: err.Error(),
			}
		}
		return &response.Response{
			Status:  http.StatusCreated,
			Data:    reqs,
			Message: "get all requests  sent by current user successfull",
		}, nil
	}
	return &response.Response{},
		utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "unauthorized user",
		}

}

func (s UserService) acceptDoctorRequuest(ctx context.Context, req *Request) (*response.Response, utilError.ApiErrorInterface) {
	reciver, err := s.UserData.GetDoctorByID(ctx, float64(req.DoctorID))
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    "line 220",
		}
	}
	fbd := &FollowedByDoctor{
		DoctorID: reciver.ID,
		UserID:   req.SenderID,
	}
	fd := &FollowingDoctor{
		DoctorID: req.SenderID,
		UserID:   reciver.ID,
	}
	_, err = s.UserData.CreateFBD(ctx, fbd)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    "line 237",
		}
	}
	_, err = s.UserData.CreateFD(ctx, fd)

	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    "line 246",
		}
	}

	return &response.Response{}, nil
}
func (s UserService) acceptPatientRequest(ctx context.Context, req *Request) (*response.Response, utilError.ApiErrorInterface) {

	reciver, err := s.UserData.GetDoctorByID(ctx, float64(req.DoctorID))
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    "line 220",
		}
	}
	fbd := &FollowedByPatient{
		DoctorID: reciver.ID,
		UserID:   req.SenderID,
	}
	fd := &FollowedDoctorsByPatient{
		PatientID: req.SenderID,
		UserID:    reciver.ID,
	}
	_, err = s.UserData.CreateFBP(ctx, fbd)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    "line 237",
		}
	}
	_, err = s.UserData.CreateFDBP(ctx, fd)

	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    "line 246",
		}
	}

	return &response.Response{}, nil

}
func (s UserService) UpdateStatusOfReq(ctx context.Context, reqId string, status string) (*response.Response, utilError.ApiErrorInterface) {
	userType := util.GetFromServiceCtx(ctx, "userType")
	if userType != DoctorClient {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "unauthorized user",
		}
	}
	req, err := s.UserData.GetReq(ctx, reqId)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Code:    "line 204",
		}
	}
	switch status {
	case "accept":
		req.Status = ACCEPTED
		_, err := s.UserData.UpdateReq(ctx, req)
		if err != nil {
			return &response.Response{}, utilError.ApiError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Code:    "line 204",
			}
		}
		if req.GenratedBy == DoctorClient {
			return s.acceptDoctorRequuest(ctx, req)
		}
		return s.acceptPatientRequest(ctx, req)
	case "reject":
		req.Status = REJECTED
		_, err := s.UserData.UpdateReq(ctx, req)
		if err != nil {
			return &response.Response{}, utilError.ApiError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Code:    "line 204",
			}
		}
	}
	return &response.Response{}, nil
}

//                                          |
//TODO: return all these with asssociation\|/
func (s UserService) GetFBD(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	fbds, err := s.UserData.GetFBD(ctx, id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status: http.StatusOK,
		Data:   fbds,
	}, nil

}
func (s UserService) GetFD(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	fbds, err := s.UserData.GetFD(ctx, id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status: http.StatusOK,
		Data:   fbds,
	}, nil
}
func (s UserService) GetFBP(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	fbds, err := s.UserData.GetFBP(ctx, id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	// patient,err:= s.UserData.GetPatientByID(ctx,)
	return &response.Response{
		Status: http.StatusOK,
		Data:   fbds,
	}, nil
}
func (s UserService) GetFDBP(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
	fbds, err := s.UserData.GetFDBP(ctx, id)
	if err != nil {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status: http.StatusOK,
		Data:   fbds,
	}, nil
}

//     ^
//   /|\
//    |

//REpport Service
func (us UserService) IsFriend(context.Context, int) (bool, utilError.ApiErrorInterface) {
	return true, nil
}
