package twtr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Errors struct {
	Status     string  `json:"status"`
	StatusCode int     `json:"status_code"`
	Errors     []Error `json:"errors"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s (error code: %d)", e.Message, e.Code)
}

func (e *Errors) Error() string {
	return fmt.Sprintf("%s: %v", e.Status, e.Errors)
}

func NewErrors(r *http.Response) *Errors {
	if r == nil {
		return &Errors{}
	} else if r.StatusCode < 400 {
		return nil
	} else if r.Body == nil {
		return &Errors{
			Status:     r.Status,
			StatusCode: r.StatusCode,
		}
	}
	defer r.Body.Close()
	var data Errors
	json.NewDecoder(r.Body).Decode(&data)
	data.Status = r.Status
	data.StatusCode = r.StatusCode
	return &data
}
