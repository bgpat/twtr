package twtr

import (
	"log"

	"github.com/garyburd/go-oauth/oauth"
)

type Client struct {
	OAuthClient  oauth.Client
	AccessToken  *oauth.Credentials
	RequestToken *oauth.Credentials
	URLFormat    string
}

func NewClient(consumer, token *oauth.Credentials) *Client {
	return &Client{
		OAuthClient: oauth.Client{
			TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
			ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
			TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
			Credentials:                   *consumer,
		},
		AccessToken: token,
		URLFormat:   "https://api.twitter.com/1.1/%s.json",
	}
}

func (c *Client) RequestTokenURL(callback string) (string, error) {
	credential, err := c.OAuthClient.RequestTemporaryCredentials(nil, callback, nil)
	if err != nil {
		return "", err
	}
	c.RequestToken = credential
	return c.OAuthClient.AuthorizationURL(c.RequestToken, nil), nil
}

func (c *Client) GetAccessToken(verifier string) error {
	token, resp, err := c.OAuthClient.RequestToken(nil, c.RequestToken, verifier)
	if err != nil {
		return err
	}
	c.RequestToken = nil
	c.AccessToken = token
	log.Printf("%#v\n", resp)
	return nil
}
