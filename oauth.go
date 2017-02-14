package twtr

import (
	"github.com/garyburd/go-oauth/oauth"
)

type Credentials struct {
	oauth.Credentials
}

func NewCredentials(token, secret string) *Credentials {
	return &Credentials{
		Credentials: oauth.Credentials{
			Token:  token,
			Secret: secret,
		},
	}
}
