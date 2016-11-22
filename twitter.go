package twitter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
)

type Client struct {
	OAuthClient   oauth.Client
	ConsumerToken *oauth.Credentials
	AccessToken   *oauth.Credentials
	RequestToken  *oauth.Credentials
	URLFormat     string
}

func NewClient(consumer, token *oauth.Credentials) *Client {
	return &Client{
		OAuthClient: oauth.Client{
			TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
			ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
			TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
			Credentials:                   *consumer,
		},
		ConsumerToken: consumer,
		AccessToken:   token,
		URLFormat:     "https://api.twitter.com/1.1/%s.json",
	}
}

func Decode(resp *http.Response, data interface{}) error {
	if err := NewErrors(resp); err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(data)
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

func (c *Client) ParseURL(path string, params url.Values) string {
	for k, v := range params {
		if len(k) <= 0 || len(v) <= 0 || k[0] != ':' {
			continue
		}
		path = strings.Replace(path, k, strings.Join(v, ","), -1)
		delete(params, k)
	}
	return fmt.Sprintf(c.URLFormat, path)
}

func (c *Client) GET(path string, params url.Values, data interface{}) error {
	uri := c.ParseURL(path, params)
	resp, err := c.OAuthClient.Get(nil, c.AccessToken, uri, params)
	if err != nil {
		return err
	}
	return Decode(resp, data)
}
