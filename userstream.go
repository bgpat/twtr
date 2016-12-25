package twtr

type UserStream struct {
	*Streaming
	Error               chan error
	Tweet               chan Tweet
	DeleteTweet         chan UserStreamDeleteTweetEvent
	FriendList          chan UserStreamFriendList
	Favorite            chan UserStreamEvent
	Unfavorite          chan UserStreamEvent
	Follow              chan UserStreamEvent
	Unfollow            chan UserStreamEvent
	ListMemberAdded     chan UserStreamEvent
	ListMemberRemoved   chan UserStreamEvent
	DirectMessage       chan UserStreamDirectMessageEvent
	DeleteDirectMessage chan UserStreamDeleteDirectMessageEvent
}

type UserStreamEvent struct {
	Event        string `json:"event"`
	Source       User   `json:"source"`
	Target       User   `json:"target"`
	TargetObject List   `json:"target_object"`
	CreatedAt    Time   `json:"created_at"`
}

type UserStreamFriendList struct {
	Friends []int64 `json:"friends"`
}

type UserStreamDeleteTweetEvent struct {
	Status      DeletedTweet `json:"status"`
	TimestampMS int64        `json:"timestamp_ms"`
}

type DeletedTweet struct {
	ID
	UserID    int64  `json:"user_id"`
	UserIDStr string `json:"user_id_str"`
}

type UserStreamDirectMessageEvent struct {
	DirectMessage DirectMessage `json:"direct_message"`
}

type UserStreamDeleteDirectMessageEvent struct {
	DirectMessage DeletedDirectMessage `json:"direct_message"`
}

type DeletedDirectMessage struct {
	ID
	UserID int64 `json:"user_id"`
}

//func (c *Client) NewUserStream(params Values) (*UserStream, error) {
func (c *Client) NewUserStream(params Values) (*Streaming, error) {
	//streaming, err := c.NewStreaming("GET", c.UserStreamURL, params)
	streaming, err := c.NewStreaming("GET", c.UserStreamURL, params)
	if err != nil {
		return nil, err
	}
	//userStream := UserStream{Streaming: streaming}
	//go userStream.worker()
	//return &userStream, nil
	return streaming, nil
}

func (u *UserStream) Close() {
	u.Streaming.Close()
	u.Streaming = nil
}

/*
func (u *UserStream) worker() {
	for u.Streaming != nil {
		hash := make(map[string]interface{})
		text, err := u.Streaming.Decode(&hash)
		if err != nil && u.Error != nil {
			u.Error <- err
			continue
		}
		if _, ok := hash["friends"]; ok && u.FriendList != nil {
			friendList := UserStreamFriendList{}
			json.NewDecoder(strings.NewReader(text)).Decode(&friendList)
			u.FriendList <- friendList
		} else if eventType, ok := hash["event"]; ok {
			switch eventType {
			case "favorite":
				event := UserStreamEvent{}
				json.NewDecoder(strings.NewReader(text)).Decode(&event)
				u.Favorite <- event
			case "unfavorite":
				event := UserStreamEvent{}
				json.NewDecoder(strings.NewReader(text)).Decode(&event)
				u.Unfavorite <- event
			case "follow":
				event := UserStreamEvent{}
				json.NewDecoder(strings.NewReader(text)).Decode(&event)
				u.Follow <- event
			case "unfollow":
				event := UserStreamEvent{}
				json.NewDecoder(strings.NewReader(text)).Decode(&event)
				u.Unfollow <- event
			case "list_member_added":
				event := UserStreamEvent{}
				json.NewDecoder(strings.NewReader(text)).Decode(&event)
				u.ListMemberAdded <- event
			case "list_member_removed":
				event := UserStreamEvent{}
				json.NewDecoder(strings.NewReader(text)).Decode(&event)
				u.ListMemberRemoved <- event
			default:
				log.Println(text)
			}
		} else if _, ok := hash["delete"]; ok && u.DeleteTweet != nil {
			event := UserStreamDeleteTweetEvent{}
			json.NewDecoder(strings.NewReader(text)).Decode(&event)
			u.DeleteTweet <- event
		} else if _, ok := hash["created_at"]; ok && u.Tweet != nil {
			tweet := Tweet{}
			json.NewDecoder(strings.NewReader(text)).Decode(&tweet)
			u.Tweet <- tweet
		} else {
			log.Println(text)
		}
	}
}
*/
