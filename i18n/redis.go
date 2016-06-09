package i18n

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/op/go-logging"
	"golang.org/x/text/language"
)

//RedisProvider redis provider
type RedisProvider struct {
	Redis  *redis.Pool     `inject:""`
	Logger *logging.Logger `inject:""`
}

//Set set locale
func (p *RedisProvider) Set(lng *language.Tag, code, message string) {
	c := p.Redis.Get()
	defer c.Close()
	if _, err := c.Do("SET", p.key(lng, code), message); err != nil {
		p.Logger.Error(err)
	}
}

//Get get locale
func (p *RedisProvider) Get(lng *language.Tag, code string) string {
	c := p.Redis.Get()
	defer c.Close()
	val, err := redis.String(c.Do("GET", p.key(lng, code)))
	if err != nil {
		p.Logger.Error(err)
		return ""
	}
	return val

}

//Del locale
func (p *RedisProvider) Del(lng *language.Tag, code string) {
	c := p.Redis.Get()
	defer c.Close()
	if _, err := c.Do("DEL", p.key(lng, code)); err != nil {
		p.Logger.Error(err)
	}
}

//Keys list all keys
func (p *RedisProvider) Keys(lng *language.Tag) ([]string, error) {
	c := p.Redis.Get()
	defer c.Close()
	return redis.Strings(c.Do("KEYS", p.key(lng, "*")))

}
func (p *RedisProvider) key(lng *language.Tag, code string) string {
	return fmt.Sprintf("locale://%s/%s", lng.String(), code)
}
