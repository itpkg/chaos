package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"gopkg.in/vmihailenco/msgpack.v2"
)

//Cache cache model
type Cache struct {
	Redis  *redis.Pool     `inject:""`
	Logger *logging.Logger `inject:""`
	Prefix string          `inject:"cache.prefix"`
}

//Keys cache items
func (p *Cache) Keys() ([]string, error) {
	c := p.Redis.Get()
	defer c.Close()
	return redis.Strings(c.Do("KEYS", fmt.Sprintf("%s://*", p.Prefix)))
}

//Flush clear cache
func (p *Cache) Flush() error {
	c := p.Redis.Get()
	defer c.Close()
	keys, err := redis.Values(c.Do("KEYS", fmt.Sprintf("%s://*", p.Prefix)))
	if err == nil && len(keys) > 0 {
		_, err = c.Do("DEL", keys...)
	}
	return err
}

//Page cache page(by url and locale)
func (p *Cache) Page(exp time.Duration, fn gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := p.key(c)
		con := p.Redis.Get()
		defer con.Close()

		if buf, err := redis.Bytes(con.Do("GET", key)); err == nil {
			var rsp CacheResponse
			if err = msgpack.Unmarshal(buf, &rsp); err == nil {
				c.Writer.WriteHeader(rsp.Status)
				c.Writer.Header().Set("Content-Type", rsp.Type)
				c.Writer.Write(rsp.Data)
				return
			}
			p.Logger.Error(err)
		} else {
			p.Logger.Error(err)
		}
		wrt := &cachedWriter{
			ResponseWriter: c.Writer,
			Exp:            exp,
			Key:            key,
			Conn:           con,
			Logger:         p.Logger,
			status:         http.StatusOK,
		}
		c.Writer = wrt
		fn(c)
	}
}

func (p *Cache) key(c *gin.Context) string {
	return fmt.Sprintf(
		"%s://%s%s",
		p.Prefix,
		c.MustGet("locale").(string),
		c.Request.URL.RequestURI(),
	)
}

//-----------------------------------------------------------------------------

//CacheResponse page cache item
type CacheResponse struct {
	Status int
	//Header http.Header
	Type string
	Data []byte
}

type cachedWriter struct {
	gin.ResponseWriter
	Conn   redis.Conn
	Key    string
	Exp    time.Duration
	Logger *logging.Logger

	written bool
	status  int
}

func (p *cachedWriter) WriteHeader(code int) {
	p.status = code
	p.written = true
	p.ResponseWriter.WriteHeader(code)
}

func (p *cachedWriter) Status() int {
	return p.status
}

func (p *cachedWriter) Written() bool {
	return p.written
}

func (p *cachedWriter) Write(data []byte) (int, error) {

	ret, err := p.ResponseWriter.Write(data)
	if err == nil {
		//cache response
		rep := CacheResponse{
			Status: p.status,
			Type:   p.Header().Get("Content-Type"),
			Data:   data,
		}
		if buf, err1 := msgpack.Marshal(rep); err1 == nil {
			if _, err = p.Conn.Do("SET", p.Key, buf, "EX", int(p.Exp/time.Second)); err != nil {
				p.Logger.Error(err)
			}
		} else {
			p.Logger.Error(err1)
		}
	}
	return ret, err
}
