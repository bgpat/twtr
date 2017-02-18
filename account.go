package twtr

type AccountSettings struct {
	AllowContributorRequest   string           `json:"allow_contributor_request"`
	AllowDMsFrom              string           `json:"allow_dms_from"`
	AllowDMGroupsFrom         string           `json:"allow_dm_groups_from"`
	AlwaysUseHttps            bool             `json:"always_use_https"`
	DiscoverableByEmail       bool             `json:"discoverable_by_email"`
	DiscoverableByMobilePhone bool             `json:"discoverable_by_mobile_phone"`
	DisplaySensitiveMedia     bool             `json:"display_sensitive_media"`
	GeoEnabled                bool             `json:"geo_enabled"`
	Language                  string           `json:"language"`
	Protected                 bool             `json:"protected"`
	ScreenName                string           `json:"screen_name"`
	ShowAllInlineMedia        bool             `json:"show_all_inline_media"`
	SleepTime                 AccountSleepTime `json:"sleep_time"`
	TimeZone                  TimeZone         `json:"time_zone"`
	TranslatorType            string           `json:"translator_type"`
	TrendLocation             []TrendLocation  `json:"trend_location"`
	UseCookiePersonalization  bool             `json:"use_cookie_personalization"`
}

type AccountSleepTime struct {
	Enabled   bool `json:"enabled"`
	EndTime   int  `json:"end_time"`
	StartTime int  `json:"start_time"`
}

func (c *Client) GetAccountSettings(params *Params) (*AccountSettings, *Response, error) {
	settings := new(AccountSettings)
	resp, err := c.GET("account/settings", params, settings)
	return settings, resp, err
}

func (c *Client) UpdateAccountSettings(params *Params) (*AccountSettings, *Response, error) {
	settings := new(AccountSettings)
	resp, err := c.POST("account/settings", params, settings)
	return settings, resp, err
}

func (c *Client) VerifyCredentials(params *Params) (*User, *Response, error) {
	user := new(User)
	resp, err := c.GET("account/verify_credentials", params, user)
	return user, resp, err
}

// UpdateProfile sets values that users are able to set under the "Account" tab of their settings page.
// https://dev.twitter.com/rest/reference/post/account/update_profile
func (c *Client) UpdateProfile(params *Params) (*User, *Response, error) {
	user := new(User)
	resp, err := c.POST("account/update_profile", params, user)
	return user, resp, err
}
