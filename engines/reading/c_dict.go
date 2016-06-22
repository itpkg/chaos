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

func (p *Engine) postDict(c *gin.Context) {
	kw := c.Query("keyword")
	if dictExp.MatchString(kw) {
		if out, err := exec.
			Command("sdcv", "--data-dir", dictROOT, kw).
			Output(); err == nil {
			c.String(http.StatusOK, string(out))
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, fmt.Sprintf("bad keyword: %s", kw))
	}

}
