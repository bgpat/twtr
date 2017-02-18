package twtr

// Cookie represents HTTP cookie
type Cookie struct {
	Path     string `json:"cookie_path"`
	Name     string `json:"cookie_name"`
	Value    string `json:"cookie_value"`
	Secure   bool   `json:"cookie_secure"`
	HTTPOnly bool   `json:"cookie_httponly"`
	Maxage   int64  `json:"cookie_maxage"`
	Domain   string `json:"cookie_domain"`
}
