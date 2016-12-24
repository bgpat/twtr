package twtr

type Entities struct {
	Hashtags     []HashtagEntity     `json:"hashtags"`
	Media        []MediaEntity       `json:"media"`
	URLs         []URLEntity         `json:"urls"`
	UserMentions []UserMentionEntity `json:"user_mentions"`
	Symbols      []SymbolEntity      `json:"symbols"`
}

type Entity struct {
	Indices [2]int `json:"indices"`
}

type HashtagEntity SymbolEntity

type MediaEntity struct {
	URLEntity
	ID
	MediaURL      string           `json:"media_url"`
	MediaURLHttps string           `json:"media_url_https"`
	Sizes         MediaEntitySizes `json:"sizes"`
	Type          string           `json:"type"`
	VideoInfo     *VideoInfo       `json:"video_info"`
}

type MediaEntitySize struct {
	H      int    `json:"h"`
	Resize string `json:"resize"`
	W      int    `json:"w"`
}

type MediaEntitySizes struct {
	Thumb  MediaEntitySize `json:"thumb"`
	Large  MediaEntitySize `json:"large"`
	Medium MediaEntitySize `json:"medium"`
	Small  MediaEntitySize `json:"small"`
}

type URLEntity struct {
	Entity
	DisplayURL  string `json:"display_url"`
	ExpandedURL string `json:"expanded_url"`
	URL         string `json:"url"`
}

type UserMentionEntity struct {
	Entity
	ID
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type SymbolEntity struct {
	Entity
	Text string `json:"text"`
}

type VideoInfo struct {
	AspectRatio    [2]int             `json:"aspect_ratio"`
	DurationMillis int                `json:"duration_millis"`
	Variants       []VideoInfoVariant `json:"variants"`
}

type VideoInfoVariant struct {
	Bitrate     int    `json:"bitrate"`
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
}
