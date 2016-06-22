package reading

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

const dictROOT = "tmp/reading/dict"

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
	if out, err := exec.
		Command("sdcv", "--data-dir", dictROOT, c.Param("keyword")).
		Output(); err == nil {
		c.String(http.StatusOK, string(out))
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}

}
