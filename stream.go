package twtr

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Stream struct {
	*http.Response
	Data chan string
}

func (c *Client) NewStream(method, urlStr string, params Values) (*Stream, error) {
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
	stream := &Stream{
		Response: resp,
		Data:     ch,
	}
	go stream.worker()
	return stream, nil
}

func (s *Stream) Close() {
	s.Response.Body.Close()
	ch := make(chan struct{}, 0)
	s.Response.Request.Cancel = ch
	close(ch)
	s.Data = nil
}

func (s *Stream) Decode(data interface{}) error {
	for s.Data != nil {
		str := <-s.Data
		if str == "" {
			continue
		}
		return json.NewDecoder(strings.NewReader(str)).Decode(data)
	}
	return errors.New("stream is not initialized")
}

func (s *Stream) worker() {
	scanner := bufio.NewScanner(s.Response.Body)
	for scanner.Scan() {
		s.Data <- scanner.Text()
	}
}
