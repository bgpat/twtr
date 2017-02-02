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

func (c *Client) GetAccountSettings(params *Values) (*AccountSettings, error) {
	settings := new(AccountSettings)
	err := c.GET("account/settings", params, settings)
	return settings, err
}

func (c *Client) UpdateAccountSettings(params *Values) (*AccountSettings, error) {
	settings := new(AccountSettings)
	err := c.POST("account/settings", params, settings)
	return settings, err
}

func (c *Client) VerifyCredentials(params *Values) (*User, error) {
	user := new(User)
	err := c.GET("account/verify_credentials", params, user)
	return user, err
}
