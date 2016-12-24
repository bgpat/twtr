package twtr

type Tweet struct {
	ID
	Contributors         []Contributor          `json:"contributors"`
	Coordinates          *Coordinates           `json:"coordinates"`
	CreatedAt            Time                   `json:"created_at"`
	CurrentUserRetweet   ID                     `json:"current_user_retweet"`
	Entities             Entities               `json:"entities"`
	ExtendedEntities     *Entities              `json:"extended_entities"`
	FavoriteCount        *int                   `json:"favorite_count"`
	Favorited            *bool                  `json:"favorited"`
	FilterLevel          string                 `json:"filter_level"`
	Geo                  *interface{}           `json:"geo"`
	InReplyToScreenName  *string                `json:"in_reply_to_screen_name"`
	InReplyToStatusID    *int64                 `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr *string                `json:"in_reply_to_status_id_str"`
	InReplyToUserID      *int64                 `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   *string                `json:"in_reply_to_user_id_str"`
	Lang                 *string                `json:"lang"`
	Place                *Place                 `json:"place"`
	PossiblySensitive    *bool                  `json:"possibly_sensitive"`
	QuotedStatusID       *int64                 `json:"quoted_status_id"`
	QuotedStatusIDStr    *string                `json:"quoted_status_id_str"`
	QuotedStatus         *Tweet                 `json:"quoted_status"`
	Scopes               map[string]interface{} `json:"scopes"`
	RetweetCount         *int                   `json:"retweet_count"`
	Retweeted            bool                   `json:"retweeted"`
	RetweetedStatus      *Tweet                 `json:"retweeted_status"`
	Source               string                 `json:"source"`
	Text                 string                 `json:"text"`
	Truncated            bool                   `json:"truncated"`
	User                 User                   `json:"user"`
	WithheldCopyright    bool                   `json:"withheld_copyright"`
	WithheldInCountries  []string               `json:"withheld_in_countries"`
	WithheldScope        string                 `json:"withheld_scope"`
}

type Contributor struct {
	ID
	ScreenName string `json:"screen_name"`
}

type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"`
	Type        string     `json:"type"`
}
