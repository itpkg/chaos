package web

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/facebookgo/inject"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/jrallison/go-workers"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func IocAction(fn func(*cli.Context, *inject.Graph) error) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		var inj inject.Graph
		logger := Logger()
		if !IsProduction() {
			inj.Logger = logger
		}

		db, err := OpenDatabase()
		if err != nil {
			return err
		}
		rep := OpenRedis()

		wfg := viper.GetStringMapString("workers.config")
		wfg["server"] = fmt.Sprintf(
			"%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"))
		wfg["database"] = viper.GetString("redis.db")
		workers.Configure(wfg)

		if err := inj.Provide(
			&inject.Object{Value: logger},
			&inject.Object{Value: db},
			&inject.Object{Value: rep},
		); err != nil {
			return err
		}
		Loop(func(en Engine) error {
			if e := en.Map(&inj); e != nil {
				return e
			}
			return inj.Provide(&inject.Object{Value: en})
		})
		if err := inj.Populate(); err != nil {
			return err
		}
		return fn(ctx, &inj)
	}
}

func Action(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		return f(c)
	}
}

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
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the app server",
			Action: IocAction(func(*cli.Context, *inject.Graph) error {
				if IsProduction() {
					gin.SetMode(gin.ReleaseMode)
				}
				rt := gin.Default()

				Loop(func(en Engine) error {
					en.Mount(rt)
					return nil
				})

				adr := fmt.Sprintf(":%d", viper.GetInt("http.port"))
				hnd := cors.New(cors.Options{
					AllowCredentials: true,
					AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
					//Debug:            true,
				}).Handler(rt)

				if IsProduction() {
					return endless.ListenAndServe(adr, hnd)
				} else {
					return http.ListenAndServe(adr, hnd)
				}
			}),
		},
		{
			Name:    "worker",
			Aliases: []string{"w"},
			Usage:   "start the worker progress",
			Action: IocAction(func(*cli.Context, *inject.Graph) error {
				Loop(func(en Engine) error {
					en.Worker()
					return nil
				})
				workers.Run()
				return nil
			}),
		},

		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "database operations",
			Subcommands: []cli.Command{
				{
					Name:    "migrate",
					Usage:   "migrate the database",
					Aliases: []string{"m"},
					Action: Action(func(*cli.Context) error {
						db, err := OpenDatabase()
						if err != nil {
							return err
						}
						return Loop(func(en Engine) error {
							en.Migrate(db)
							return nil
						})
					}),
				},
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
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("redis", map[string]interface{}{
		"host": "localhost",
		"port": 6379,
		"db":   2,
	})
	viper.SetDefault("workers.config", map[string]interface{}{
		"pool":      30,
		"namespace": "tasks",
		"process":   RandomStr(8),
	})
}
