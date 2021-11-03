package asset

import (
	utilError "askUs/v1/util/error"
	"askUs/v1/util/response"
	"context"
	"io"
	"net/http"

	"cloud.google.com/go/storage"
)

type (
	Service interface {
		Upload(context.Context, *UploadRequest) (*response.Response, utilError.ApiErrorInterface)
		Download(context.Context, string, http.ResponseWriter) (*response.Response, utilError.ApiErrorInterface)
	}
	Store interface {
		New() *storage.BucketHandle
		Post(context.Context, *storage.BucketHandle, string, io.Reader) error
		Get(context.Context, *storage.BucketHandle, string) ([]byte, error)
	}
	UploadRequest struct {
		FileName string
		Kind     string
		Reader   io.Reader
	}
	UploadResponse struct {
		Url string `json:"url"`
	}
)

const (
	Profile = "profile"
	Advice  = "advice"
	Report  = "resport"
)
