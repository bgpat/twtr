package twtr

import (
	"log"

	"github.com/garyburd/go-oauth/oauth"
)

type Client struct {
	OAuthClient   oauth.Client
	AccessToken   *Credentials
	RequestToken  *Credentials
	URLFormat     string
	UserStreamURL string
}

func NewClient(consumer, token *Credentials) *Client {
	return &Client{
		OAuthClient: oauth.Client{
			TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
			ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
			TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
			Credentials:                   consumer.Credentials,
		},
		AccessToken:   token,
		URLFormat:     "https://api.twitter.com/1.1/%s.json",
		UserStreamURL: "https://userstream.twitter.com/1.1/user.json",
	}
}

func (c *Client) RequestTokenURL(callback string) (string, error) {
	credential, err := c.OAuthClient.RequestTemporaryCredentials(nil, callback, nil)
	if err != nil {
		return "", err
	}
	c.RequestToken = &Credentials{Credentials: *credential}
	return c.OAuthClient.AuthorizationURL(credential, nil), nil
}

func (c *Client) GetAccessToken(verifier string) error {
	token, resp, err := c.OAuthClient.RequestToken(nil, &c.RequestToken.Credentials, verifier)
	if err != nil {
		return err
	}
	c.RequestToken = nil
	c.AccessToken = &Credentials{Credentials: *token}
	log.Printf("%#v\n", resp)
	return nil
}
