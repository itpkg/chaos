package web

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/unrolled/render"
)

type Controller struct {
	Render *render.Render  `inject:""`
	Db     *gorm.DB        `inject:""`
	Logger *logging.Logger `inject:""`
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
