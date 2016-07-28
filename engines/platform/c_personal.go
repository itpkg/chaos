package platform

import "net/http"

func (p *Engine) deleteSignOut(wrt http.ResponseWriter, req *http.Request) (interface{}, error) {

	u, e := p.Jwt.CurrentUser(req)
	if e != nil {
		return nil, e
	}
	p.Dao.Log(u.ID, "sign out")
	return p.Ok(), nil
}

func (p *Engine) getPersonalLogs(wrt http.ResponseWriter, req *http.Request) (interface{}, error) {
	u, e := p.Jwt.CurrentUser(req)
	if e != nil {
		return nil, e
	}
	var logs []Log
	err := p.Dao.Db.
		Select([]string{"created_at", "message"}).
		Where("user_id = ?", u.ID).
		Order("ID DESC").Limit(200).
		Find(&logs).Error
	return logs, err
}

func (p *Engine) getPersonalSelf(wrt http.ResponseWriter, req *http.Request) (interface{}, error) {
	return p.Jwt.CurrentUser(req)
}
