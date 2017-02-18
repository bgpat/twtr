# twtr

[![Build Status](https://travis-ci.org/bgpat/twtr.svg?branch=master)](https://travis-ci.org/bgpat/twtr)
[![Go Report Card](https://goreportcard.com/badge/github.com/bgpat/twtr)](https://goreportcard.com/report/github.com/bgpat/twtr)
[![codecov](https://codecov.io/gh/bgpat/twtr/branch/master/graph/badge.svg)](https://codecov.io/gh/bgpat/twtr)
[![GoDoc](https://godoc.org/github.com/bgpat/twtr?status.svg)](http://godoc.org/github.com/bgpat/twtr)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

## Installation

Install and update with `go get github.com/bgpat/twtr`

## Examples

### REST API

```go
consumer := twtr.NewCredentials(consumerKey, consumerSecret)
token := twtr.NewCredentials(accessToken, accessSecret)
client := twtr.NewClient(consumer, token)

// account/verify_credentials
user, resp, err := client.VerifyCredentials(nil)
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

	user, _, err := client.GetUser(&twtr.Params{"screen_name": "Twitter"})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("user:", user)
```
