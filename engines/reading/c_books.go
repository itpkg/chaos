package reading

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/chonglou/epubgo"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
)

const booksROOT = "tmp/reading/books"

func (p *Engine) indexBooks(c *gin.Context) (interface{}, error) {
	var books []Book
	err := p.Db.Order("id ASC").Find(&books).Error
	return books, err
}

func (p *Engine) showBook(c *gin.Context) {

	name := c.Param("name")

	var book Book
	if err := p.Db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if buf, err := p._readBook(&book, name); err == nil {
		web.Bytes(name, buf)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}

}

//-----------------------------------------------------------------------------
func (p *Engine) _readBook(book *Book, name string) ([]byte, error) {
	bk, err := epubgo.Open(fmt.Sprintf("%s/%s", booksROOT, book.Name))
	if err != nil {
		return nil, err
	}

	for it, err := bk.Spine(); !it.IsLast(); it.Next() {
		if err != nil {
			return nil, err
		}
		p.Logger.Debugf("url: %s, name: %s", it.URL(), name)
		if name[1:] == it.URL() {
			pg, err := it.Open()
			defer pg.Close()
			if err != nil {
				return nil, err
			}
			return ioutil.ReadAll(pg)
		}
	}

	return nil, errors.New("not found")
}

func (p *Engine) _scanBooks() error {
	const sep = ","
	const ext = ".epub"

	var books []Book
	err := filepath.Walk(booksROOT, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() && filepath.Ext(info.Name()) == ext {
			p.Logger.Infof("find book %s", path)
			bk, err := epubgo.Open(path)
			if err != nil {
				return err
			}
			title, _ := bk.Metadata("title")
			subject, _ := bk.Metadata("subject")
			publisher, _ := bk.Metadata("publisher")
			creator, _ := bk.Metadata("creator")
			version, _ := bk.Metadata("date")

			book := Book{
				Name:      path[len(booksROOT)+1 : len(path)],
				Title:     strings.Join(title, sep),
				Subject:   strings.Join(subject, sep),
				Publisher: strings.Join(publisher, sep),
				Creator:   strings.Join(creator, sep),
				Version:   strings.Join(version, sep),
			}

			for it, err := bk.Navigation(); !it.IsLast(); err = it.Next() {
				if err != nil {
					return err
				}
				for _, u := range []string{"TableOfContents.xhtml"} {
					if strings.HasPrefix(it.URL(), u) {
						book.Home = u
						break
					}
				}
				if book.Home != "" {
					break
				}
			}
			if book.Home == "" {
				return fmt.Errorf("Bad book format: %+v", book)
			}

			books = append(books, book)
		}
		return nil
	})

	if err == nil {
		for _, b := range books {
			var bk Book
			if err = p.Db.Where("name = ?", b.Name).Find(&bk).Error; err == nil {
				err = p.Db.Model(&bk).Updates(b).Error
			} else {
				err = p.Db.Create(&b).Error
			}
			if err != nil {
				break
			}
		}

	}

	return err
}
