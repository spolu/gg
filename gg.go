package main

import (
	"flag"
	"fmt"
	"os"

	"./account"
	"./oauth"
)

func init() {

}

func main() {
	// gg -create -a stripe
	// gg -a stripe -request_url
	// gg -a stripe -oauth_code xxx

	var accountName string
	const (
		aDefault = "default"
		aUsage   = "The account to use"
	)
	flag.StringVar(&accountName, "a", aDefault, aUsage+" (shorthand)")

	var oauthCode string
	const (
		oaDefault = ""
		oaUsage   = "The code to input after OAuth authorization"
	)
	flag.StringVar(&oauthCode, "oauth_code", oaDefault, oaUsage)

	var authorizeURL bool
	const (
		auDefault = false
		auUsage   = "Dumps the authorization URL"
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

	account := account.NewAccount(accountName)

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
