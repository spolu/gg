package main

import (
	"flag"
	"fmt"
	"os"

	"./oauth"
)

func init() {

}

func main() {
	// gg -create -a stripe
	// gg -a stripe -request_url
	// gg -a stripe -oauth_code xxx

	var account string
	const (
		aDefault = "default"
		aUsage   = "The account to use"
	)
	flag.StringVar(&account, "a", aDefault, aUsage+" (shorthand)")

	var oauthCode string
	const (
		oaDefault = ""
		oaUsage   = "The code to input after OAuth authorization"
	)
	flag.StringVar(&oauthCode, "oauth_code", oaDefault, oaUsage)

	var authorizeURL bool
	const (
		auDefault = false
		auUsage   = "The code to input after OAuth authorization"
	)
	flag.BoolVar(&authorizeURL, "authorize_url", auDefault, auUsage)

	flag.Parse()

	if authorizeURL {
		oa := new(oauth.OAuth)
		u, err := oa.GetAuthorizeURL()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", u)
		os.Exit(0)
	}

	if oauthCode != "" {
		oa := new(oauth.OAuth)
		_, err := oa.RetrieveRefreshToken(oauthCode)
		if err != nil {
			panic(err)
		}
		// TODO(spolu) store refresh token in account configuration
		os.Exit(0)
	}

	/*
		g := gmail.NewGmail(OAUTH2_REFRESH_TOKEN)
		g.Do()
	*/
}
