package twtr

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dustin/go-jsonpointer"
)

type Streaming struct {
	*http.Response
	Event chan interface{}
	Error chan error
	text  chan string
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
	WithheldInCountries []string `json:"withheld_in_countires"`
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

func (c *Client) NewStreaming(method, urlStr string, params Values) (*Streaming, error) {
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
	ch := make(chan string)
	streaming := &Streaming{
		Response: resp,
		text:     ch,
	}
	go streaming.polling()
	return streaming, nil
}

func (s *Streaming) Close() {
	s.Response.Body.Close()
	ch := make(chan struct{})
	s.Response.Request.Cancel = ch
	close(ch)
	s.text = nil
}

func (s *Streaming) Decode() (interface{}, error) {
	for s.text != nil {
		json := []byte(<-s.text)
		if e := new(StreamingTweetEvent); validEventType(e, json, "/target_object/source") {
			return e, nil
		} else if e := new(StreamingListEvent); validEventType(e, json, "/target_object/slug") {
			return e, nil
		} else if e := new(StreamingGeneralEvent); validEventType(e, json, "/event") {
			return e, nil
		} else if e := new(Tweet); validEventType(e, json, "/source") {
			return e, nil
		} else if e := new(StreamingDirectMessageEvent); validEventType(e, json, "/direct_message") {
			return e, nil
		} else if e := new(StreamingDeleteTweetEvent); validEventType(e, json, "/delete/status") {
			return e, nil
		} else if e := new(StreamingDeleteDirectMessageEvent); validEventType(e, json, "/delete/direct_message") {
			return e, nil
		} else if e := new(StreamingDeleteLocationEvent); validEventType(e, json, "/scrub_geo") {
			return e, nil
		} else if e := new(StreamingLimitEvent); validEventType(e, json, "/limit") {
			return e, nil
		} else if e := new(StreamingStatusWithheldEvent); validEventType(e, json, "/status_withheld") {
			return e, nil
		} else if e := new(StreamingUserWithheldEvent); validEventType(e, json, "/user_withheld") {
			return e, nil
		} else if e := new(StreamingDisconnectEvent); validEventType(e, json, "/disconnect") {
			return e, nil
		} else if e := new(StreamingWarningEvent); validEventType(e, json, "/warning") {
			return e, nil
		} else if e := new(StreamingEnvelopeEvent); validEventType(e, json, "/for_user") {
			return e, nil
		} else if e := new(StreamingControlEvent); validEventType(e, json, "/control") {
			return e, nil
		} else if e := new(StreamingFriendsEvent); validEventType(e, json, "/friends") {
			return e, nil
		}
	}
	return nil, errors.New("streaming is not yet initialized")
}

func (s *Streaming) polling() {
	scanner := bufio.NewScanner(s.Response.Body)
	for scanner.Scan() {
		s.text <- scanner.Text()
		//if s.Event != nil {
		event, err := s.Decode()
		if err == nil {
			s.Event <- event
		} else if s.Error != nil {
			s.Error <- err
		}
		//}
	}
}

func validEventType(event interface{}, data []byte, path string) bool {
	if v, _ := jsonpointer.Find(data, path); v == nil {
		return false
	}
	err := json.Unmarshal(data, event)
	return err == nil
}
