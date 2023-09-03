package helper

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func ReadRequestBody(request *http.Request, result interface{}) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func WriteResponse(writer http.ResponseWriter, code int, stauts string, responseBody interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	encoder := json.NewEncoder(writer)
	response := JsonResponse{
		Code:   code,
		Status: stauts,
		Data:   responseBody,
	}

	encoder.Encode(response)
}
