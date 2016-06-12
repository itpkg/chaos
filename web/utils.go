package web

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

//IsProduction production mode ?
func IsProduction() bool {
	return viper.GetString("env") == "production"
}

//Host hostname
func Host() string {
	if IsProduction() {
		if viper.GetBool("http.ssl") {
			return fmt.Sprintf("https://%s", viper.GetString("http.domain"))
		}
		return fmt.Sprintf("http://%s", viper.GetString("http.domain"))
	}
	return fmt.Sprintf("http://localhost:%d", viper.GetInt("http.port"))
}

//RandomStr random string
func RandomStr(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	buf := make([]rune, n)
	for i := range buf {
		buf[i] = letters[rand.Intn(len(letters))]
	}
	return string(buf)
}

//OpenDatabase open database
func OpenDatabase() (*gorm.DB, error) {
	//postgresql: "user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
	args := ""
	for k, v := range viper.GetStringMapString("database.args") {
		args += fmt.Sprintf(" %s=%s ", k, v)
	}
	db, err := gorm.Open(viper.GetString("database.driver"), args)
	if err != nil {
		return nil, err
	}
	if !IsProduction() {
		db.LogMode(true)
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(viper.GetInt("database.pool.max_idle"))
	db.DB().SetMaxOpenConns(viper.GetInt("database.pool.max_open"))
	return db, nil

}

//OpenRedis open redis
func OpenRedis() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, e := redis.Dial(
				"tcp",
				fmt.Sprintf(
					"%s:%d",
					viper.GetString("redis.host"),
					viper.GetInt("redis.port"),
				),
			)
			if e != nil {
				return nil, e
			}
			if _, e = c.Do("SELECT", viper.GetInt("redis.db")); e != nil {
				c.Close()
				return nil, e
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
