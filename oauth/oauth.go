package oauth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	CLIENT_ID       = "711658480682-fk78b5n8h6n35u74ketg7dn61v1f6ktu.apps.googleusercontent.com"
	CLIENT_SECRET   = "c6lc6xaM956QAetkANU2DEPe"
	REDIRECT_URI    = "urn:ietf:wg:oauth:2.0:oob"
	OAUTH_AUTH_URL  = "https://accounts.google.com/o/oauth2/auth"
	OAUTH_TOKEN_URL = "https://accounts.google.com/o/oauth2/token"
	OAUTH_SCOPES    = "https://www.googleapis.com/auth/gmail.modify"
)

type OAuth struct{}

func (oa *OAuth) GetAuthorizeURL() (*url.URL, error) {
	v := url.Values{}
	v.Add("response_type", "code")
	v.Add("client_id", CLIENT_ID)
	v.Add("redirect_uri", REDIRECT_URI)
	v.Add("scope", OAUTH_SCOPES)
	v.Add("access_type", "offline")

	return url.Parse(OAUTH_AUTH_URL + "?" + v.Encode())
}

func (oa *OAuth) RetrieveRefreshToken(code string) (string, error) {
	v := url.Values{}
	v.Add("code", code)
	v.Add("client_id", CLIENT_ID)
	v.Add("client_secret", CLIENT_SECRET)
	v.Add("redirect_uri", REDIRECT_URI)
	v.Add("grant_type", "authorization_code")

	res, err := http.PostForm(OAUTH_TOKEN_URL, v)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var dat map[string]interface{}
	err = json.Unmarshal(body, &dat)
	if err != nil {
		return "", err
	}

	if e, ok := dat["error"].(string); ok {
		return "", errors.New(e)
	}

	if refreshToken, ok := dat["refresh_token"].(string); ok {
		return refreshToken, nil
	} else {
		return "", errors.New("No `refresh_token` returned")
	}
}
