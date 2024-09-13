package errors

import (
	"encoding/json"
	"net/http"
)

type BizErr struct {
	Tag     string `json:"tag"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func (e *BizErr) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func GeneralBizErr(statusCode int, cause string) *BizErr {
	return &BizErr{
		Tag:     tags[statusCode],
		Message: msgs[statusCode],
		Cause:   cause,
	}
}

var tags = map[int]string{
	http.StatusBadRequest:          "BAD_REQUEST",
	http.StatusUnauthorized:        "UNAUTHORIZED",
	http.StatusForbidden:           "FORBIDDEN",
	http.StatusNotFound:            "NOT_FOUND",
	http.StatusConflict:            "CONFLICT",
	http.StatusInternalServerError: "INTERNAL_SERVER_ERROR",
}

var msgs = map[int]string{
	http.StatusBadRequest:          "Invalid request.",
	http.StatusUnauthorized:        "Not authorized.",
	http.StatusForbidden:           "No permission to access.",
	http.StatusNotFound:            "Resource not found.",
	http.StatusConflict:            "Resource already exists.",
	http.StatusInternalServerError: "Server unavailable.",
}
