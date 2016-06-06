package platform

import (
	"github.com/itpkg/chaos/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "database operations",
			Subcommands: []cli.Command{
				{
					Name:    "migrate",
					Usage:   "migrate the database",
					Aliases: []string{"m"},
					Action: web.Action(func(*cli.Context) error {
						db, err := web.OpenDatabase()
						if err != nil {
							return err
						}
						return web.Loop(func(en web.Engine) error {
							en.Migrate(db)
							return nil
						})
					}),
				},
			},
		},
	}
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
	viper.SetDefault("workers.email", 5)
}
