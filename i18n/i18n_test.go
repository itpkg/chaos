package i18n_test

import (
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/itpkg/chaos/i18n"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/op/go-logging"
	"golang.org/x/text/language"
)

var logger = logging.MustGetLogger("test")
var lang = &language.SimplifiedChinese

func TestDatabase(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal(err)
	}
	db.LogMode(true)
	i18n.Migrate(db)
	testProvider(t, &i18n.DatabaseProvider{Db: db, Logger: logger})
}

func TestRedis(t *testing.T) {
	re := redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				"localhost:6379",
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	testProvider(t, &i18n.RedisProvider{Redis: &re, Logger: logger})

}

func testProvider(t *testing.T, p i18n.Provider) {
	key := "hello"
	val := "你好"
	p.Set(lang, key, val)
	p.Set(lang, key+".1", val)
	if val1 := p.Get(lang, key); val != val1 {
		t.Errorf("want %s, get %s", val, val1)
	}
	ks, err := p.Keys(lang)
	if err != nil {
		t.Fatal(err)
	}
	if len(ks) == 0 {
		t.Errorf("empty keys")
	} else {
		t.Log(ks)
	}
	p.Del(lang, key)
}
