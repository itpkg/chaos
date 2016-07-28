package web

import (
	"net/http"
	"time"

	"golang.org/x/text/language"

	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/unrolled/render"
)

type Controller struct {
	Render *render.Render  `inject:""`
	Db     *gorm.DB        `inject:""`
	Logger *logging.Logger `inject:""`
}

func (p *Controller) Ok() map[string]interface{} {
	return map[string]interface{}{"ok": true, "created": time.Now()}
}

func (p *Controller) Locale(req *http.Request) *language.Tag {
	// 1. Check URL arguments.
	lng := req.URL.Query().Get("locale")

	// 2. Get language information from cookies.
	if len(lng) == 0 {
		if ck, er := req.Cookie("locale"); er == nil {
			lng = ck.Value
		}
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lng) == 0 {
		al := req.Header.Get("Accept-Language")
		if len(al) > 4 {
			lng = al[:5]
		}
	}

	tag, _, _ := matcher.Match(language.Make(lng))
	// 	c.SetCookie("locale", tag.String(), 1<<31-1, "/", "", false, false)
	return &tag

}

func (p *Controller) Json(w http.ResponseWriter, v interface{}, e error) {
	if e == nil {
		p.Render.JSON(w, http.StatusOK, v)
	} else {
		p.Render.Text(w, http.StatusInternalServerError, e.Error())
	}
}

// //SendBytes write bytes
// func (p *Controller) WriteBytes(w http.ResponseWriter, name string, buf []byte) {
//
// 	ext := path.Ext(name)
// 	switch {
// 	case ext == ".md" || ext == ".txt":
// 		p.Render.Text(w, http.StatusOK, string(buf))
// 	case ext == ".html" || ext == ".htm":
// 		p.Render.HTML(w, status, name, binding, htmlOpt)
// 		c.Data(http.StatusOK, "text/html; charset=UTF-8", buf)
// 	case ext == ".xhtml":
// 		c.Data(http.StatusOK, "application/xhtml+xml; charset=UTF-8", buf)
// 	case ext == ".xml":
// 		c.Data(http.StatusOK, "text/xml; charset=UTF-8", buf)
// 	case ext == ".css":
// 		c.Data(http.StatusOK, "text/css; charset=UTF-8", buf)
// 	case ext == ".png" || ext == ".jpg" || ext == ".jpeg":
// 		c.Data(http.StatusOK, fmt.Sprintf("image/%s", ext), buf)
// 	default:
// 		c.String(http.StatusNotFound, fmt.Sprintf("bad format %s", ext))
// 	}
//
// }

//=============================================================================

var matcher language.Matcher

func init() {
	matcher = language.NewMatcher([]language.Tag{
		language.AmericanEnglish,
		language.SimplifiedChinese,
	})
}
