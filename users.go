package twitter

type User struct {
	ProfileBackgroundTile          bool   `json:"profile_background_tile,omitempty"`
	ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color,omitempty"`
	Name                           string `json:"name,omitempty"`
	ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color,omitempty"`
	Location                       string `json:"location,omitempty"`
	CreatedAt                      Time   `json:"created_at,omitempty"`
	ProfileImageURL                string `json:"profile_image_url,omitempty"`
	IsTranslator                   bool   `json:"is_translator,omitempty"`
	IDStr                          string `json:"id_str,omitempty"`
	ProfileLinkColor               string `json:"profile_link_color,omitempty"`
	FollowRequestSent              bool   `json:"follow_request_sent,omitempty"`
	FavouritesCount                int    `json:"favourites_count,omitempty"`
	ContributorsEnabled            bool   `json:"contributors_enabled,omitempty"`
	URL                            string `json:"url,omitempty"`
	DefaultProfile                 bool   `json:"default_profile,omitempty"`
	ProfileImageURLHttps           string `json:"profile_image_url_https,omitempty"`
	UTCOffset                      int    `json:"utc_offset,omitempty"`
	ProfileBannerURL               string `json:"profile_banner_url,omitempty"`
	ID                             int64  `json:"id,omitempty"`
	ProfileUseBackgroundImage      bool   `json:"profile_use_background_image,omitempty"`
	ListedCount                    int    `json:"listed_count,omitempty"`
	ProfileTextColor               string `json:"profile_text_color,omitempty"`
	Protected                      bool   `json:"protected,omitempty"`
	Lang                           string `json:"lang,omitempty"`
	FollowersCount                 int    `json:"followers_count,omitempty"`
	Description                    string `json:"description,omitempty"`
	Notifications                  bool   `json:"notifications,omitempty"`
	GeoEnabled                     bool   `json:"geo_enabled,omitempty"`
	Verified                       bool   `json:"verified,omitempty"`
	TimeZone                       string `json:"time_zone,omitempty"`
	ProfileBackgroundColor         string `json:"profile_background_color,omitempty"`
	ProfileBackgroundImageURLHttps string `json:"profile_background_image_url_https,omitempty"`
	StatusesCount                  int    `json:"statuses_count,omitempty"`
	FriendsCount                   int    `json:"friends_count,omitempty"`
	DefaultProfileImage            bool   `json:"default_profile_image,omitempty"`
	ProfileBackgroundImageURL      string `json:"profile_background_image_url,omitempty"`
	Following                      bool   `json:"following,omitempty"`
	ScreenName                     string `json:"screen_name,omitempty"`
}
