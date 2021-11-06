package report

import (
	"askUs/v1/package/asset"
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"
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
		Create(context.Context, *UserReport) (*UserReport, error)
		Update(context.Context, *UserReport) (*UserReport, error)
		Get(context.Context, string, float64) (*UserReport, error)
		Delete(context.Context, *UserReport) (*UserReport, error)
		GetAll(context.Context, int) (*[]UserReport, error)
	}

	Asset interface {
		Upload(context.Context, *asset.UploadRequest) (*response.Response, utilError.ApiErrorInterface)
		Download(context.Context, string, http.ResponseWriter) (*response.Response, utilError.ApiErrorInterface)
		Delete(ctx context.Context, fileurl string) utilError.ApiErrorInterface
	}

	User interface {
		// A service layer to check given id of patient and current user are friends or not
		IsFriend(context.Context, int) (bool, utilError.ApiErrorInterface)
	}

	Service interface {
		Create(context.Context, *asset.UploadRequest, string) (*response.Response, utilError.ApiErrorInterface)
		Update(context.Context, string, *asset.UploadRequest, string) (*response.Response, utilError.ApiErrorInterface)
		Delete(context.Context, string) (*response.Response, utilError.ApiErrorInterface)
		GetAll(context.Context, int) (*response.Response, utilError.ApiErrorInterface)
	}
	UserReport struct {
		ID        int    `json:"id"  gorm:"primaryKey"`
		Url       string `json:"url"`
		Header    string `json:"header"`
		FileUrl   string `json:"file_url"`
		PatientId int    `json:"patent_id"`
	}
)

const (
	Report = "report"
)

func (UserReport) TableName() string {
	return "report.reports"
}
