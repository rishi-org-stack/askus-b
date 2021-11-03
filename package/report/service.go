package report

import (
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"
	"io"
	"net/http"
)

//dependendence on asset and user module
//api's post /report/add fom value header is new (header defines the what this report is for)
///report/update/resportid -put basically deletes old file and creates newfile with same name but different  file
//report/delete/reportid -delete basically deletes that file
//resport/ get gets all report
//report/get/report/;patientID- get cross check with user module both are connected or not
type (
	DB interface {
		Create(context.Context, *UserReport) (*UserReport, utilError.ApiErrorInterface)
		Update(context.Context, *UserReport) (*UserReport, utilError.ApiErrorInterface)
		Delete(context.Context, *UserReport) (*UserReport, utilError.ApiErrorInterface)
		GetAll(context.Context, int) (*[]UserReport, utilError.ApiErrorInterface)
	}

	Asset interface {
		Upload(context.Context, *UploadRequest) (*response.Response, utilError.ApiErrorInterface)
		Download(context.Context, string, http.ResponseWriter) (*response.Response, utilError.ApiErrorInterface)
	}

	User interface {
		// A service layer to check given id of patient and current user are friends or not
		IsFriend(context.Context, int) (bool, utilError.ApiErrorInterface)
	}

	Service interface {
		Create(context.Context, *UploadRequest) (*UserReport, utilError.ApiErrorInterface)
		Update(context.Context, int) (*UserReport, utilError.ApiErrorInterface)
		Delete(context.Context, int) (*UserReport, utilError.ApiErrorInterface)
		GetAll(context.Context, int) (*[]UserReport, utilError.ApiErrorInterface)
	}
	UserReport struct {
		ID        int    `json:"id"  gorm:"primaryKey"`
		Url       string `json:"url"`
		Header    string `json:"header"`
		PatientId int    `json:"patent_id"`
	}
	UploadRequest struct {
		FileName string
		// Kind     string
		Reader io.Reader
	}
	UploadResponse struct {
		Url string `json:"url"`
	}
)

const (
	Report = "resport"
)

func (UserReport) TableName() string {
	return "report.reports."
}
