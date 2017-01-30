package twtr

type User struct {
	ID
	ContributorsEnabled            bool     `json:"contributors_enabled"`
	CreatedAt                      Time     `json:"created_at"`
	DefaultProfile                 bool     `json:"default_profile"`
	DefaultProfileImage            bool     `json:"default_profile_image"`
	Description                    *string  `json:"description"`
	Entities                       Entities `json:"entities"`
	FavouritesCount                int      `json:"favourites_count"`
	FollowRequestSent              *bool    `json:"follow_request_sent"`
	Following                      *bool    `json:"following"`
	FollowersCount                 int      `json:"followers_count"`
	FriendsCount                   int      `json:"friends_count"`
	GeoEnabled                     bool     `json:"geo_enabled"`
	IsTranslator                   bool     `json:"is_translator"`
	Lang                           string   `json:"lang"`
	ListedCount                    int      `json:"listed_count"`
	Location                       *string  `json:"location"`
	Name                           string   `json:"name"`
	Notifications                  bool     `json:"notifications"`
	ProfileBackgroundColor         string   `json:"profile_background_color"`
	ProfileBackgroundImageURL      string   `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHttps string   `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool     `json:"profile_background_tile"`
	ProfileBannerURL               string   `json:"profile_banner_url"`
	ProfileImageURL                string   `json:"profile_image_url"`
	ProfileImageURLHttps           string   `json:"profile_image_url_https"`
	ProfileLinkColor               string   `json:"profile_link_color"`
	ProfileSidebarBorderColor      string   `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string   `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string   `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool     `json:"profile_use_background_image"`
	Protected                      bool     `json:"protected"`
	ScreenName                     string   `json:"screen_name"`
	ShowAllInlineMedia             bool     `json:"show_all_inline_media"`
	Status                         *Tweet   `json:"status"`
	StatusesCount                  int      `json:"statuses_count"`
	TimeZone                       *string  `json:"time_zone"`
	URL                            *string  `json:"url"`
	UTCOffset                      *int     `json:"utc_offset"`
	Verified                       bool     `json:"verified"`
	WithheldInCountries            string   `json:"withheld_in_countries"`
	WithheldScope                  string   `json:"withheld_scope"`
}

type UserIDs struct {
	Cursor
	IDs []int64 `json:"ids"`
}

func (c *Client) GetUser(params *Values) (*User, error) {
	user := new(User)
	err := c.GET("users/show", params, user)
	return user, err
}

func (c *Client) GetUsers(params *Values) ([]User, error) {
	var users []User
	err := c.GET("users/lookup", params, &users)
	return users, err
}

func (c *Client) GetUserProfileBanner(params *Values) (*Images, error) {
	images := new(Images)
	err := c.GET("users/profile_banner", params, images)
	return images, err
}

func (c *Client) SearchUsers(params *Values) ([]User, error) {
	var users []User
	err := c.GET("users/search", params, &users)
	return users, err
}

func (c *Client) ReportSpamUser(params *Values) (*User, error) {
	user := new(User)
	err := c.POST("users/report_spam", params, user)
	return user, err
}

func (c *Client) GetMyProfile(params *Values) (*User, error) {
	user := new(User)
	err := c.GET("account/verify_credentials", params, user)
	return user, err
}
