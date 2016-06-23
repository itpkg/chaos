package reading

import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"
)

const dictROOT = "tmp/reading/dict"

var dictExp = regexp.MustCompile(`^[\p{Han}\w]{1,32}$`)

func (p *Engine) getDict(c *gin.Context) {
	if out, err := exec.
		Command("sdcv", "--data-dir", dictROOT, "-l").
		Output(); err == nil {
		c.String(http.StatusOK, string(out))
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}

}

//DictFm form for notice
type DictFm struct {
	Keyword string `form:"keyword" binding:"required"`
}

func (p *Engine) postDict(c *gin.Context) {
	var fm DictFm
	err := c.Bind(&fm)
	if err == nil {
		if dictExp.MatchString(fm.Keyword) {
			var out []byte
			if out, err = exec.
				Command("sdcv", "--data-dir", dictROOT, fm.Keyword).
				Output(); err == nil {
				c.String(http.StatusOK, string(out))
				return
			}
		} else {
			err = fmt.Errorf("bad keyword: %s", fm.Keyword)
		}
	}
	c.String(http.StatusInternalServerError, err.Error())
}
