package twtr

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/dustin/go-jsonpointer"
)

type Streaming struct {
	*http.Response
	Event   chan interface{}
	Error   chan error
	scanner *bufio.Scanner
	listen  *sync.Mutex
}

type StreamingGeneralEvent struct {
	Event        string      `json:"event"`
	Source       User        `json:"source"`
	Target       User        `json:"target"`
	TargetObject interface{} `json:"target_object"`
	CreatedAt    Time        `json:"created_at"`
}

type StreamingTweetEvent struct {
	StreamingGeneralEvent
	TargetObject Tweet `json:"target_object"`
}

type StreamingListEvent struct {
	StreamingGeneralEvent
	TargetObject List `json:"target_object"`
}

type StreamingDirectMessageEvent struct {
	DirectMessage DirectMessage `json:"direct_message"`
}

type StreamingDeleteTweetEvent struct {
	Delete struct {
		Status      StreamingDeletedObject `json:"status"`
		TimestampMS int64                  `json:"timestamp_ms"`
	} `json:"delete"`
}

type StreamingDeleteDirectMessageEvent struct {
	Delete struct {
		DirectMessage StreamingDeletedObject `json:"direct_message"`
	} `json:"delete"`
}

type StreamingDeletedObject struct {
	ID
	UserID    int64  `json:"user_id"`
	UserIDStr string `json:"user_id_str"`
}

type StreamingDeleteLocationEvent struct {
	ScrubGeo struct {
		UserID          int64  `json:"user_id"`
		UserIDStr       string `json:"user_id_str"`
		UpToStatusID    int64  `json:"up_to_status_id"`
		UpToStatusIDStr string `json:"up_to_status_id_str"`
	} `json:"scrub_geo"`
}

type StreamingLimitEvent struct {
	Limit struct {
		Track int `json:"track"`
	} `json:"limit"`
}

type StreamingStatusWithheldEvent struct {
	StatusWithheld StreamingWithheldObject `json:"status_withheld"`
}

type StreamingUserWithheldEvent struct {
	UserWithheld StreamingWithheldObject `json:"user_withheld"`
}

type StreamingWithheldObject struct {
	ID                  int64    `json:"id"`
	UserID              int64    `json:"user_id"`
	WithheldInCountries []string `json:"withheld_in_countries"`
}

type StreamingDisconnectEvent struct {
	Disconnect struct {
		Code       int    `json:"code"`
		StreamName string `json:"stream_name"`
		Reason     string `json:"reason"`
	} `json:"disconnect"`
}

type StreamingWarningEvent struct {
	Warning struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		PercentFull int    `json:"percent_full"`
		UserID      int64  `json:"user_id"`
	} `json:"warning"`
}

type StreamingEnvelopeEvent struct {
	ForUser int64 `json:"for_user"`
	Message struct {
		Friends []int64 `json:"friends"`
	} `json:"message"`
}

type StreamingControlEvent struct {
	Control struct {
		ControlURI string `json:"control_uri"`
	} `json:"control"`
}

type StreamingFriendsEvent struct {
	Friends []int64 `json:"friends"`
}

func (c *Client) NewStreaming(method, urlStr string, params *Params) (*Streaming, error) {
	var resp *http.Response
	var err error
	if method == http.MethodGet {
		resp, err = c.getRequest(urlStr, params)
	} else if method == http.MethodPost {
		resp, err = c.postRequest(urlStr, params)
	} else {
		return nil, errors.New(method + " method is not yet supported")
	}
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(resp.Body)
	listen := new(sync.Mutex)
	streaming := &Streaming{
		Response: resp,
		scanner:  scanner,
		listen:   listen,
	}
	streaming.Stop()
	go streaming.polling()
	return streaming, nil
}

func (s *Streaming) Close() {
	s.listen = nil
	ch := make(chan struct{})
	s.Response.Request.Cancel = ch
	close(ch)
	s.Response.Body.Close()
	s.scanner = nil
}

func (s *Streaming) Start() {
	if s.listen != nil {
		s.listen.Unlock()
	}
}

func (s *Streaming) Stop() {
	if s.listen != nil {
		s.listen.Lock()
	}
}

func (s *Streaming) Decode() (interface{}, error) {
	for s.scanner != nil && s.scanner.Scan() {
		data := s.scanner.Bytes()
		if len(data) == 0 {
			continue
		}
		if e := new(StreamingTweetEvent); validEventType(e, data, "/target_object/source") {
			return e, nil
		} else if e := new(StreamingListEvent); validEventType(e, data, "/target_object/slug") {
			return e, nil
		} else if e := new(StreamingGeneralEvent); validEventType(e, data, "/event") {
			return e, nil
		} else if e := new(Tweet); validEventType(e, data, "/source") {
			return e, nil
		} else if e := new(StreamingDirectMessageEvent); validEventType(e, data, "/direct_message") {
			return e, nil
		} else if e := new(StreamingDeleteTweetEvent); validEventType(e, data, "/delete/status") {
			return e, nil
		} else if e := new(StreamingDeleteDirectMessageEvent); validEventType(e, data, "/delete/direct_message") {
			return e, nil
		} else if e := new(StreamingDeleteLocationEvent); validEventType(e, data, "/scrub_geo") {
			return e, nil
		} else if e := new(StreamingLimitEvent); validEventType(e, data, "/limit") {
			return e, nil
		} else if e := new(StreamingStatusWithheldEvent); validEventType(e, data, "/status_withheld") {
			return e, nil
		} else if e := new(StreamingUserWithheldEvent); validEventType(e, data, "/user_withheld") {
			return e, nil
		} else if e := new(StreamingDisconnectEvent); validEventType(e, data, "/disconnect") {
			return e, nil
		} else if e := new(StreamingWarningEvent); validEventType(e, data, "/warning") {
			return e, nil
		} else if e := new(StreamingEnvelopeEvent); validEventType(e, data, "/for_user") {
			return e, nil
		} else if e := new(StreamingControlEvent); validEventType(e, data, "/control") {
			return e, nil
		} else if e := new(StreamingFriendsEvent); validEventType(e, data, "/friends") {
			return e, nil
		} else {
			return nil, fmt.Errorf("failed to decode: %#v", data)
		}
	}
	return nil, errors.New("streaming is not yet initialized")
}

func (s *Streaming) polling() {
	for s.listen != nil {
		s.listen.Lock()
		s.listen.Unlock()
		event, err := s.Decode()
		if err != nil {
			if s.Error != nil {
				s.Error <- err
			}
		} else {
			if s.Event != nil {
				s.Event <- event
			}
		}
	}
}

func validEventType(event interface{}, data []byte, path string) bool {
	if v, _ := jsonpointer.Find(data, path); v == nil {
		return false
	}
	err := json.Unmarshal(data, event)
	return err == nil
}
