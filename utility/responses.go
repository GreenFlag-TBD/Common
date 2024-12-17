package utility

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewErrorResponse(message string, status int) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Status:  status,
	}
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(message string, status int, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
}
