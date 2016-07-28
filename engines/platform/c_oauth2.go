package platform

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/schema"
	"golang.org/x/oauth2"
)

//OauthFm oauth2 model
type OauthFm struct {
	Code  string `schema:"code"`
	State string `schema:"state"`
}

func (p *Engine) postOauth2Callback(wrt http.ResponseWriter, req *http.Request) (interface{}, error) {
	var fm OauthFm
	req.ParseForm()

	dec := schema.NewDecoder()
	if err := dec.Decode(&fm, req.PostForm); err != nil {
		return nil, err
	}
	//p.Logger.Debugf("%+v", fm)
	var u *User
	var e error
	switch fm.State {
	case p.Oauth2GoogleState:
		u, e = p._google(fm.Code)
	default:
		e = errors.New("bad state")
	}

	var tk []byte
	if e == nil {
		p.Dao.Log(u.ID, "sign in")
		tk, e = p.Jwt.Sum(p.Dao.UserClaims(u), 7) //TODO days
	}
	return map[string]interface{}{"token": string(tk)}, e

}

func (p *Engine) _google(code string) (*User, error) {
	var cfg oauth2.Config
	if err := p.Dao.Get("google.oauth", &cfg); err != nil {
		return nil, err
	}

	tok, err := cfg.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	cli := cfg.Client(oauth2.NoContext, tok)
	res, err := cli.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var gu GoogleUser
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&gu); err != nil {
		return nil, err
	}
	return p.Dao.AddUser(gu.ID, "google", gu.Email, gu.Name, gu.Link, gu.Picture)
}
