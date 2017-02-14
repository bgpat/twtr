# twtr

## USAGE

### REST API

```go
consumer := twtr.NewCredentials(consumerKey, consumerSecret)
token := twtr.NewCredentials(accessToken, accessSecret)
client := twtr.NewClient(consumer, token)

// account/verify_credentials
user, err := client.VerifyCredentials(nil)
```

### OAuth Verify URL

```go
consumer := twtr.NewCredentials(consumerKey, consumerSecret)
client := twtr.NewClient(consumer, nil)

url, err := client.RequestTokenURL("https://example.com/")
```
