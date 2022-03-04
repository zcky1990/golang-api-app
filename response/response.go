package response

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Status string     `json:"status"`
	Code   uint8      `json:"code"`
	Data   *fiber.Map `json:"data"`
}

type SuccessResponseData struct {
	Status string `json:"status"`
	Code   uint8  `json:"code"`
	Data   map[string]string
}

type FailedResponse struct {
	Status string `json:"status"`
	Code   uint16 `json:"code"`
	Data   map[string]string
}

func GetSuccessResponseData() SuccessResponseData {
	return SuccessResponseData{"success", 200, map[string]string{}}
}

func GetSuccessResponse(data *fiber.Map) SuccessResponse {
	return SuccessResponse{"success", 200, data}
}

func GetFailedResponse(err string) FailedResponse {
	errors := FailedResponse{"failed", 500, map[string]string{}}
	errors.Data["error_message"] = err
	return errors
}
