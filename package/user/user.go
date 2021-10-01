package user

import (
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"

	// "askUs/v1/util/response"
	"context"
	"net/http"
)

const (
	source                    = "USER"
	source_doctor             = "DOCTOR"
	source_patient            = "PATIENT"
	ID_DECODE_ERROR           = source + "_ERROR_GEETING_ID"
	USER_GET_ERROR            = source + "_GET_ERROR"
	USER_DOCTOR_CREATE_ERROR  = source + "_DOCTOR_CREATE_ERROR"
	USER_PATIENT_CREATE_ERROR = source + "_PATIENT_CREATE_ERROR"
	// USER_GET_IDEA_ERROR       = source + "_" + source_idea + "_GET_ERROR"
	// USER_AUTH_GET_ERROR       = source + "_" + source_auth + "_GET_ERROR"
	USER_COPY_ERROR = source + "_COPY_ERROR"
	// USER_IDEA_INSERT_ERROR    = source + "_" + source_idea + "_INSERT_ERROR"
	USER_PATIENT_UPDATE_ERROR = source + "_" + source_patient + "_UPDATE_ERROR"
	USER_DOCTOR_UPDATE_ERROR  = source + "_" + source_doctor + "_UPDATE_ERROR"
)

type (
	UserService struct {
		UserData DB
		// AuthService auth.Service
	}
)

func Init(db DB) Service {
	return &UserService{
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
func (s UserService) GetPatientByID(ctx context.Context) (*response.Response, utilError.ApiErrorInterface) {
	id := ctx.Value("surround").(map[string]interface{})["id"].(float64)
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
			Status:  http.StatusBadRequest,
			Code:    USER_PATIENT_UPDATE_ERROR,
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
			Status:  http.StatusBadRequest,
			Code:    USER_PATIENT_UPDATE_ERROR,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "User Update success full",
		Data:    doc,
	}, nil
}
