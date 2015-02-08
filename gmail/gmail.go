package gmail

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	GG_CLIENT_ID     = "711658480682-fk78b5n8h6n35u74ketg7dn61v1f6ktu.apps.googleusercontent.com"
	GG_CLIENT_SECRET = "c6lc6xaM956QAetkANU2DEPe"
	OAUTH2_TOKEN_URL = "https://accounts.google.com/o/oauth2/token"
)

type Gmail struct {
	refreshToken      string
	accessToken       string
	accessTokenExpiry int64
}

func NewGmail(refreshToken string) *Gmail {
	g := new(Gmail)
	g.refreshToken = refreshToken
	g.accessToken = ""
	g.accessTokenExpiry = 0
	return g
}

func (g *Gmail) refreshAccessToken() {
	now := time.Now().UnixNano() / 1000000

	if len(g.accessToken) > 0 &&
		g.accessTokenExpiry > now+1000*60*15 {
		return
	}

	res, err := http.PostForm(OAUTH2_TOKEN_URL,
		url.Values{
			"refresh_token": {g.refreshToken},
			"client_id":     {GG_CLIENT_ID},
			"client_secret": {GG_CLIENT_SECRET},
			"grant_type":    {"refresh_token"},
		})
	defer res.Body.Close()

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s", body)
}

func (g *Gmail) doGet(url *url.URL) {
	g.refreshAccessToken()
}

func (g *Gmail) Do() {
	g.refreshAccessToken()
}
