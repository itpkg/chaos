package reading

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/chaos/web"
	"github.com/itpkg/epub"
)

const booksROOT = "tmp/reading/books"

func (p *Engine) indexBooks(c *gin.Context) (interface{}, error) {
	var books []Book
	err := p.Db.Order("id ASC").Find(&books).Error
	return books, err
}

func (p *Engine) deleteBook(c *gin.Context) (interface{}, error) {
	err := p.Db.Where("id = ?", c.Param("id")).Delete(&Book{}).Error
	return web.OK, err
}

func (p *Engine) indexBook(c *gin.Context) {
	bk, err := p.book(c)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer bk.Close()
	c.JSON(http.StatusOK, bk)
}

func (p *Engine) showBook(c *gin.Context) {
	bk, err := p.book(c)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer bk.Close()

	name := c.Param("name")[1:]
	fd, err := bk.Open(name)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer fd.Close()
	buf, err := ioutil.ReadAll(fd)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	web.Bytes(name, buf)(c)

}

func (p *Engine) book(c *gin.Context) (*epub.Book, error) {
	var book Book
	if err := p.Db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return nil, err
	}
	return epub.Open(fmt.Sprintf("%s/%s", booksROOT, book.Name))

}

//-----------------------------------------------------------------------------

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
			bk, err := epub.Open(path)
			if err != nil {
				return err
			}
			defer bk.Close()

			var version []string
			for _, d := range bk.Opf.Metadata.Date {
				version = append(version, d.Data)
			}
			var creator []string
			for _, c := range bk.Opf.Metadata.Creator {
				creator = append(creator, c.Data)
			}

			book := Book{
				Name:      path[len(booksROOT)+1 : len(path)],
				Title:     strings.Join(bk.Opf.Metadata.Title, sep),
				Subject:   strings.Join(bk.Opf.Metadata.Subject, sep),
				Publisher: strings.Join(bk.Opf.Metadata.Publisher, sep),
				Creator:   strings.Join(creator, sep),
				Version:   strings.Join(version, sep),
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
