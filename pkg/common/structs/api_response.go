package structs

type APIResponse struct {
	StatusCode int64       `json:"statuscode"`
	Data       interface{} `json:"data"`
}

func (e *APIResponse) GetStatusCode() int64 {
	return e.StatusCode
}

func (e *APIResponse) GetData() interface{} {
	return e.Data
}

func (e *APIResponse) GetPtr() *APIResponse {
	return e
}

func NewAPIResponse(data interface{}, code int64) APIResponse {

	response := APIResponse{
		StatusCode: code,
		Data:       data,
	}

	return response
}
