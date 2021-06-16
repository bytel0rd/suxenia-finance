package structs

type APIResponse struct {
	StatusCode int         `json:"statuscode"`
	Data       interface{} `json:"data"`
}

func (e *APIResponse) GetStatusCode() int {
	return e.StatusCode
}

func (e *APIResponse) GetData() interface{} {
	return e.Data
}

func (e *APIResponse) GetPtr() *APIResponse {
	return e
}

func NewAPIResponse(data interface{}, code int) APIResponse {

	response := APIResponse{
		StatusCode: code,
		Data:       data,
	}

	return response
}
