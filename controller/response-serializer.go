package controller

type SuccessResponse struct {
	Status string `json:"status"`
	Code   uint8  `json:"code"`
	Data   map[string]string
}

type FailedResponse struct {
	Status string `json:"status"`
	Code   uint16 `json:"code"`
	Data   map[string]string
}

func GetSuccessResponse() SuccessResponse {
	return SuccessResponse{"success", 200, map[string]string{}}
}

func GetFailedResponse(err string) FailedResponse {
	errors := FailedResponse{"failed", 500, map[string]string{}}
	errors.Data["error_message"] = err
	return errors
}
