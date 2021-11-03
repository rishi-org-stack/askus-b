package asset

import (
	"askUs/v1/util"
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type AssetService struct {
	store Store
}

func Init(s Store) Service {
	return &AssetService{
		store: s,
	}
}

func (as AssetService) Upload(
	ctx context.Context,
	upr *UploadRequest) (*response.Response, utilError.ApiErrorInterface) {
	id := util.GetFromServiceCtx(ctx, "id").(float64)
	if upr.Kind != Profile && upr.Kind != Advice {
		return &response.Response{}, utilError.ApiError{
			Status:  http.StatusBadRequest,
			Message: "unkwon kind",
		}
	}
	fileUrl := upr.Kind + "/" + strconv.Itoa(int(id)) + "_" + upr.FileName
	reader := upr.Reader
	buc := as.store.New()
	err := as.store.Post(ctx, buc, fileUrl, reader)
	if err != nil && !strings.HasPrefix(err.Error(), "googleapi: Error 404: No such object") {
		return &response.Response{},
			utilError.ApiError{
				Status:  http.StatusBadRequest,
				Message: "failed upload :" + err.Error(),
			}
	}
	return &response.Response{
		Status:  http.StatusOK,
		Message: "successful upload",
		Data: UploadResponse{
			Url: "locolhost:8080/asset/down/" + fileUrl,
		},
	}, nil
}
func (as AssetService) Download(ctx context.Context, url string, writer http.ResponseWriter) (*response.Response, utilError.ApiErrorInterface) {
	bucket := as.store.New()
	data, err := as.store.Get(ctx, bucket, url)
	if err != nil {
		return &response.Response{},
			utilError.ApiError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
	}
	_, err = io.Copy(writer, bytes.NewReader(data))
	if err != nil {
		return &response.Response{},
			utilError.ApiError{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
	}
	return &response.Response{}, nil
}
