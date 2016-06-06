package web

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func Run() error {
	app := cli.NewApp()
	app.Name = "chaos"
	app.Version = "v20160606"
	app.Usage = "it-package web application."
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init config file",
			Action: func(*cli.Context) error {
				const fn = "config.toml"
				if _, err := os.Stat(fn); err == nil {
					return fmt.Errorf("file %s already exists!", fn)
				}

				args := viper.AllSettings()
				fd, err := os.Create(fn)
				if err != nil {
					log.Println(err)
					return err
				}
				defer fd.Close()
				end := toml.NewEncoder(fd)
				err = end.Encode(args)

				return err

			},
		},
	}
	for _, en := range engines {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}
	return app.Run(os.Args)
}

func IsProduction() bool {
	return viper.GetString("env") == "production"
}

func init() {
	viper.SetEnvPrefix("chaos")
	viper.BindEnv("env")
	viper.SetDefault("env", "development")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
}
