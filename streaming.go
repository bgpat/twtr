package twtr

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Streaming struct {
	*http.Response
	Data chan string
}

func (c *Client) NewStreaming(method, urlStr string, params Values) (*Streaming, error) {
	var resp *http.Response
	var err error
	if method == "GET" {
		resp, err = c.getRequest(urlStr, params)
	} else if method == "POST" {
		resp, err = c.postRequest(urlStr, params)
	} else {
		return nil, errors.New(method + " is not implemented method")
	}
	if err != nil {
		return nil, err
	}
	ch := make(chan string)
	streaming := &Streaming{
		Response: resp,
		Data:     ch,
	}
	go streaming.worker()
	return streaming, nil
}

func (s *Streaming) Close() {
	s.Response.Body.Close()
	ch := make(chan struct{})
	s.Response.Request.Cancel = ch
	close(ch)
	s.Data = nil
}

func (s *Streaming) Decode(data interface{}) (string, error) {
	for s.Data != nil {
		text := <-s.Data
		if text == "" {
			continue
		}
		return text, json.NewDecoder(strings.NewReader(text)).Decode(data)
	}
	return "", errors.New("streaming is not initialized")
}

func (s *Streaming) worker() {
	scanner := bufio.NewScanner(s.Response.Body)
	for scanner.Scan() {
		s.Data <- scanner.Text()
	}
}
