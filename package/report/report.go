package report

import (
	"askUs/v1/package/asset"
	"askUs/v1/util"
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"
	"net/http"
	"strings"
)

type ReportService struct {
	db           DB
	userService  User
	assetService Asset
}

func Init(db DB, us User, as Asset) Service {
	return &ReportService{
		db:           db,
		userService:  us,
		assetService: as,
	}
}

func (rs ReportService) Create(ctx context.Context, ur *asset.UploadRequest, header string) (
	*response.Response,
	utilError.ApiErrorInterface) {

	id := util.GetFromServiceCtx(ctx, "id").(float64)
	ur.Kind = Report
	res, err := rs.assetService.Upload(ctx, ur)
	if err != nil {
		return nil, err
	}
	uploasRes := res.Data.(asset.UploadResponse)
	fileurl := uploasRes.Url[strings.Index(uploasRes.Url, "report/"):]
	urr, err := rs.db.Create(ctx, &UserReport{
		Url:       uploasRes.Url,
		Header:    header,
		FileUrl:   fileurl,
		PatientId: int(id),
	})
	if err != nil {
		return nil, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status:  http.StatusCreated,
		Data:    urr,
		Message: "successfully created report",
	}, nil
}

func (rs ReportService) Update(
	ctx context.Context,
	rid string,
	upr *asset.UploadRequest,
	header string) (
	*response.Response,
	utilError.ApiErrorInterface) {

	id := util.GetFromServiceCtx(ctx, "id").(float64)
	ur, err := rs.db.Get(ctx, rid, id)
	if err != nil {
		return nil, &utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	err = rs.assetService.Delete(ctx, ur.FileUrl)
	if err != nil {
		return nil, &utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	upr.Kind = Report
	res, err := rs.assetService.Upload(ctx, upr)
	uploadRes := res.Data.(asset.UploadResponse)
	if err != nil {
		return nil, err
	}
	fileurl := uploadRes.Url[strings.Index(uploadRes.Url, "report/"):]

	ur.Header = header
	ur.FileUrl = fileurl
	ur.Url = uploadRes.Url

	ur, err = rs.db.Update(ctx, ur)

	if err != nil {
		return nil, &utilError.ApiError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "update a userReport successfl",
		Data:    ur,
	}, nil
}
func (rs ReportService) Delete(ctx context.Context, rid string) (
	*response.Response,
	utilError.ApiErrorInterface) {
	id := util.GetFromServiceCtx(ctx, "id").(float64)
	ur, err := rs.db.Get(ctx, rid, id)
	if err != nil {
		return nil, &utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	err = rs.assetService.Delete(ctx, ur.FileUrl)
	if err != nil {
		return nil, err
	}

	_, err = rs.db.Delete(ctx, ur)

	if err != nil {
		return nil, &utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "delete a userReport successfl",
		// Data:    ur,
	}, nil
}
func (rs ReportService) GetAll(ctx context.Context, rid int) (
	*response.Response,
	utilError.ApiErrorInterface) {
	return nil, nil
}
