package platform

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

//NoticeFm form for notice
type NoticeFm struct {
	Content string `schema:"content" validate:"required"`
}

func (p *Engine) getNotices(wrt http.ResponseWriter, req *http.Request) (interface{}, error) {
	var ns []Notice
	err := p.Dao.Db.
		Select([]string{"id", "content", "created_at"}).
		Order("ID DESC").Limit(120).Find(&ns).Error
	return ns, err
}

func (p *Engine) postNotices(wrt http.ResponseWriter, req *http.Request) (interface{}, error) {
	lng := p.Locale(req)
	req.ParseForm()

	var fm NoticeFm
	dec := schema.NewDecoder()
	if err := dec.Decode(&fm, req.PostForm); err != nil {
		return nil, err
	}
	n := Notice{Content: fm.Content, Lang: lng.String()}
	err := p.Dao.Db.Create(&n).Error
	return n, err
}

func (p *Engine) deleteNotice(wrt http.ResponseWriter, req *http.Request) (interface{}, error) {
	id := mux.Vars(req)["id"]
	err := p.Dao.Db.Where("id = ?", id).Delete(&Notice{}).Error
	return p.Ok(), err
}
