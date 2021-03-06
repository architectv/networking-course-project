package models

type ApiResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *ApiResponse) Set(code int, message string, data interface{}) {
	r.Code = code
	r.Message = message
	r.Data = data
}

func (r *ApiResponse) Error(code int, message string) {
	r.Code = code
	r.Message = message
}

// func Send(code int, message string, data interface{}) *ApiResponse {
// 	return &ApiResponse{
// 		Code:    code,
// 		Message: message,
// 		Data:    data,
// 	}
// }

// func Error(code int, message string) *ApiResponse {
// 	return &ApiResponse{
// 		Code:    code,
// 		Message: message,
// 	}
// }
