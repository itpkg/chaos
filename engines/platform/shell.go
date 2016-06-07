package platform

import (
	"github.com/itpkg/chaos/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	viper.SetDefault("http", map[string]interface{}{
		"port":   8080,
		"domain": "localhost",
		"ssl":    false,
	})
	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"user":    "postgres",
			"dbname":  "chaos",
			"sslmode": "disable",
		},
		"pool": map[string]int{
			"max_open": 180,
			"max_idle": 6,
		},
	})
	viper.SetDefault("secrets", web.RandomStr(512))
	viper.SetDefault("workers.queues.email", 5)
}
