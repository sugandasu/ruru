package tolo

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

func ResponseJson(w http.ResponseWriter, code int, data any) {
	res, _ := json.Marshal(data)
	w.WriteHeader(code)
	w.Write(res)
}

func ResponseSuccess(w http.ResponseWriter, message string, data any) {
	res := Response{
		Data:    data,
		Message: message,
	}

	resJson, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}

func ResponseError(w http.ResponseWriter, err error) {
	apperror := ParseError(err)

	res := Response{
		Message: err.Error(),
		Errors:  err.Error(),
	}

	code := http.StatusBadRequest
	if apperror != nil {
		res.Errors = apperror.Data()
		code = apperror.StatusCode()
	}

	resJson, _ := json.Marshal(res)

	w.WriteHeader(code)
	w.Write(resJson)
}
