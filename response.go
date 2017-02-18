package twtr

import (
	"errors"
	"net/http"
	"strconv"
)

type Response struct {
	*http.Response
	*Limit
}

func (r *Response) ParseResponse(resp *http.Response) error {
	r.Response = resp
	return r.parseLimit()
}

func (r *Response) parseLimit() error {
	l := Limit{}

	limit, err := r.getHeader("X-Rate-Limit-Limit")
	if err != nil {
		return err
	}
	l.Limit, err = strconv.Atoi(limit)
	if err != nil {
		return err
	}

	remaining, err := r.getHeader("X-Rate-Limit-Remaining")
	if err != nil {
		return err
	}
	l.Remaining, err = strconv.Atoi(remaining)
	if err != nil {
		return err
	}

	reset, err := r.getHeader("X-Rate-Limit-Reset")
	if err != nil {
		return err
	}
	err = l.Reset.ParseUnixTime(reset)
	if err != nil {
		return err
	}

	r.Limit = &l
	return nil
}

func (r *Response) getHeader(name string) (string, error) {
	if r.Response.Header == nil {
		return "", errors.New("header is nil")
	}
	v, ok := r.Response.Header[name]
	if !ok {
		return "", errors.New(name + " header is undefined")
	}
	if len(v) < 1 {
		return "", errors.New(name + "header is empty")
	}
	return v[0], nil
}
