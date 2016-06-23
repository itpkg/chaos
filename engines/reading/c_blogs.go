package reading

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (p *Engine) indexBlogs(c *gin.Context) (interface{}, error) {
	return p._scanBlogs()
}

func (p *Engine) showBlog(c *gin.Context) {
	name := c.Param("name")
	if buf, err := ioutil.ReadFile(fmt.Sprintf("%s%s", blogsRoot, name)); err == nil {
		c.String(http.StatusOK, string(buf))
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

//-----------------------------------------------------------------------------
const blogsRoot = "tmp/reading/blogs"

func (p *Engine) _scanBlogs() (map[string]string, error) {
	blogs := make(map[string]string)
	const ext = ".md"
	err := filepath.Walk(blogsRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() && filepath.Ext(info.Name()) == ext {
			fd, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fd.Close()
			sc := bufio.NewScanner(fd)
			if sc.Scan() {
				blogs[path[len(blogsRoot)+1:]] = sc.Text()
			}
			return sc.Err()
		}
		return nil
	})
	return blogs, err
}
