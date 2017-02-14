# twtr

## Installation

Install and update with `go get github.com/bgpat/twtr`

## Examples

### REST API

```go
consumer := twtr.NewCredentials(consumerKey, consumerSecret)
token := twtr.NewCredentials(accessToken, accessSecret)
client := twtr.NewClient(consumer, token)

// account/verify_credentials
user, err := client.VerifyCredentials(nil)
```

### Get OAuth Token

```go
	consumer := twtr.NewCredentials(consumerKey, consumerSecret)
	client := twtr.NewClient(consumer, nil)

	url, _ := client.RequestTokenURL("https://401.jp/")
	fmt.Println(url)

	var verifier string
	fmt.Scan(&verifier)
	client.GetAccessToken(verifier)

	user, _ := client.VerifyCredentials(nil)
	fmt.Println(user)
```
