package error

type ApiError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
type ApiErrorInterface interface {
	Error() string
}

func (apiError ApiError) Error() string {
	return apiError.Message
}

// func (apiError *ApiError) SourceService() string {
// 	return apiError.Source
// }
// func (apiError *ApiError) SourceServiceLevel() string {
// 	return apiError.Level
// }
