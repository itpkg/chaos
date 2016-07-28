package web

import (
	"crypto/x509/pkix"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/facebookgo/inject"
	"github.com/fvbock/endless"
	"github.com/gorilla/mux"
	"github.com/itpkg/chaos/i18n"
	"github.com/jrallison/go-workers"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

//IocAction ioc action
func IocAction(fn func(*cli.Context, *inject.Graph) error) cli.ActionFunc {
	return Action(func(ctx *cli.Context) error {
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

		i1n := i18n.I18n{Locales: make(map[string]map[string]string)}
		if err := inj.Provide(
			&inject.Object{Value: logger},
			&inject.Object{Value: db},
			&inject.Object{Value: rep},
			&inject.Object{Value: &i18n.DatabaseProvider{}},
			&inject.Object{Value: &i1n},
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
		if err := i1n.Load("locales"); err != nil {
			return err
		}
		return fn(ctx, &inj)
	})
}

//Action cfg action
func Action(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		return f(c)
	}
}

//Run main entry
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "environment, e",
					Value: "development",
					Usage: "environment, like: development, production, test...",
				},
			},
			Action: func(c *cli.Context) error {
				const fn = "config.toml"
				if _, err := os.Stat(fn); err == nil {
					return fmt.Errorf("file %s already exists", fn)
				}
				fmt.Printf("generate file %s\n", fn)

				viper.Set("env", c.String("environment"))
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
				rt := mux.NewRouter()
				Loop(func(en Engine) error {
					en.Mount(rt)
					return nil
				})

				adr := fmt.Sprintf(":%d", viper.GetInt("http.port"))
				hnd := cors.New(cors.Options{
					AllowCredentials: true,
					AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
					AllowedHeaders:   []string{"*"},
					//Debug:          true,
				}).Handler(rt)

				if IsProduction() {
					return endless.ListenAndServe(adr, hnd)
				}
				return http.ListenAndServe(adr, hnd)

			}),
		},
		{
			Name:    "routes",
			Aliases: []string{"rt"},
			Usage:   "list http routes",
			Action: func(*cli.Context) error {
				rt := mux.NewRouter()
				Loop(func(en Engine) error {
					en.Mount(rt)
					return nil
				})

				fmt.Println("HOST\tPATH\tFUNC")
				return rt.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
					hst, _ := route.GetHostTemplate()
					pth, _ := route.GetPathTemplate()
					fmt.Printf("%s\t%s\t%s\n", hst, pth, runtime.FuncForPC(reflect.ValueOf(route.GetHandler).Pointer()).Name())
					return nil
				})
			},
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
			Name:    "openssl",
			Aliases: []string{"ssl"},
			Usage:   "generate ssl certificates",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "name",
				},
				cli.StringFlag{
					Name:  "country, c",
					Value: "Earth",
					Usage: "country",
				},
				cli.StringFlag{
					Name:  "organization, o",
					Value: "Mother Nature",
					Usage: "organization",
				},
				cli.IntFlag{
					Name:  "years, y",
					Value: 1,
					Usage: "years",
				},
			},
			Action: Action(func(c *cli.Context) error {
				name := c.String("name")
				if len(name) == 0 {
					cli.ShowCommandHelp(c, "openssl")
					return nil
				}
				root := path.Join("etc", "ssl", name)

				key, crt, err := CreateCertificate(
					true,
					pkix.Name{
						Country:      []string{c.String("country")},
						Organization: []string{c.String("organization")},
					},
					c.Int("years"),
				)
				if err != nil {
					return err
				}

				fnk := path.Join(root, "key.pem")
				fnc := path.Join(root, "crt.pem")

				fmt.Printf("generate pem file %s\n", fnk)
				err = WritePemFile(fnk, "RSA PRIVATE KEY", key)
				fmt.Printf("test: openssl rsa -noout -text -in %s\n", fnk)

				if err == nil {
					fmt.Printf("generate pem file %s\n", fnc)
					err = WritePemFile(fnc, "CERTIFICATE", crt)
					fmt.Printf("test: openssl x509 -noout -text -in %s\n", fnc)
				}
				if err == nil {
					fmt.Printf(
						"verify: diff <(openssl rsa -noout -modulus -in %s) <(openssl x509 -noout -modulus -in %s)",
						fnk,
						fnc,
					)
				}
				return err

			}),
		},
	}
	for _, en := range engines {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}
	return app.Run(os.Args)
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
		"db":   8,
	})

	viper.SetDefault("http", map[string]interface{}{
		"port":   8080,
		"domain": "localhost",
		"ssl":    false,
	})
	viper.SetDefault("database", map[string]interface{}{
		"driver": "postgres",
		"args": map[string]interface{}{
			"host":    "localhost",
			"port":    5432,
			"user":    "postgres",
			"dbname":  "chaos",
			"sslmode": "disable",
		},
		"pool": map[string]int{
			"max_open": 180,
			"max_idle": 6,
		},
	})
	viper.SetDefault("secrets", RandomStr(512))

	viper.SetDefault("workers.config", map[string]interface{}{
		"pool":      30,
		"namespace": "tasks",
		"process":   RandomStr(8),
	})
}
