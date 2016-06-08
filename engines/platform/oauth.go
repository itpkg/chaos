package platform

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

/**
https://developers.google.com/identity/protocols/OAuth2WebServer
https://developers.google.com/identity/protocols/googlescopes
*/
type GoogleUser struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Link    string `json:"link"`
	Picture string `json:"picture"`
}

type GoogleCredential struct {
	Web struct {
		ClientID     string   `json:"client_id"`
		ClientSecret string   `json:"client_secret"`
		RedirectURLS []string `json:"redirect_uris"`
	} `json:"web"`
}

func (p *GoogleCredential) To() oauth2.Config {
	return oauth2.Config{
		ClientID:     p.Web.ClientID,
		ClientSecret: p.Web.ClientSecret,
		RedirectURL:  p.Web.RedirectURLS[0],
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}
