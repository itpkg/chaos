package platform

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/itpkg/chaos/web"
)

func (p *Engine) getAdminSiteInfo(wrt http.ResponseWriter, req *http.Request) {
	ifo := p._siteInfoMap(p.Locale(req))
	p.Render.JSON(wrt, http.StatusOK, ifo)
}

func (p *Engine) postAdminSiteInfo(wrt http.ResponseWriter, req *http.Request) (interface{}, error) {
	req.ParseForm()
	lng := p.Locale(req)
	for _, k := range []string{"title", "subTitle", "keywords", "description",
		"aboutUs", "copyright"} {
		if err := p.Dao.Set(p._siteKey(lng, k), req.Form.Get(k), false); err != nil {
			return nil, err
		}
	}
	for _, k := range []string{"name", "email"} {
		if err := p.Dao.Set(p._siteAuthorKey(k), req.Form.Get("author"+strings.Title(k)), false); err != nil {
			return nil, err
		}
	}
	var links []web.Link
	if err := json.Unmarshal([]byte(req.Form.Get("navLinks")), &links); err != nil {
		return nil, err
	}
	if err := p.Dao.Set(p._siteKey(nil, "navLinks"), links, false); err != nil {
		return nil, err
	}
	return p.Ok(), nil
}

func (p *Engine) deleteAdminCache(wrt http.ResponseWriter, req *http.Request) (interface{}, error) {
	err := p.Cache.Flush()
	return p.Ok(), err
}
