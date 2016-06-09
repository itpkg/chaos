package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/template"

	"github.com/facebookgo/inject"
	"github.com/itpkg/chaos/web"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "nginx",
			Aliases: []string{"ng"},
			Usage:   "init nginx config file",
			Action: web.Action(func(*cli.Context) error {
				const tpl = `
upstream {{.Domain}}_prod {
  server localhost:{{.Port}} fail_timeout=0;
}

server {
  listen {{if .Ssl}}443{{- else}}80{{- end}};

{{if .Ssl}}
  ssl  on;
  ssl_certificate  ssl/{{.Domain}}-cert.pem;
  ssl_certificate_key  ssl/{{.Domain}}-key.pem;
  ssl_session_timeout  5m;
  ssl_protocols  SSLv2 SSLv3 TLSv1;
  ssl_ciphers  RC4:HIGH:!aNULL:!MD5;
  ssl_prefer_server_ciphers  on;
{{- end}}

  client_max_body_size 4G;
  keepalive_timeout 10;

  proxy_buffers 16 64k;
  proxy_buffer_size 128k;

  server_name www.{{.Domain}} {{.Domain}};

  root {{.Root}}/public;
  index index.html;

  access_log /var/log/nginx/{{.Domain}}.access.log;
  error_log /var/log/nginx/{{.Domain}}.error.log;

  location / {
    try_files $uri $uri/ /index.html?/$request_uri;
  }

  location ^~ /assets/ {
    gzip_static on;
    expires max;
    access_log off;
    add_header Cache-Control "public";
  }

  location ~* \.(?:rss|atom)$ {
    expires 12h;
    add_header Cache-Control "public";
  }

  location ~ ^/api/{{.Version}}(/?)(.*) {
    {{if .Ssl}}proxy_set_header X-Forwarded-Proto https;{{- end}}
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_redirect off;
    proxy_pass http://{{.Domain}}_prod/$2$is_args$args;
    # limit_req zone=one;
  }

}

`
				t := template.Must(template.New("").Parse(tpl))
				pwd, err := os.Getwd()
				if err != nil {
					return err
				}
				fd, err := os.OpenFile("nginx.conf", os.O_WRONLY|os.O_CREATE, 0600)
				if err != nil {
					return err
				}
				defer fd.Close()

				return t.Execute(fd, struct {
					Domain  string
					Port    int
					Ssl     bool
					Root    string
					Version string
				}{
					Ssl:     viper.GetBool("http.ssl"),
					Domain:  viper.GetString("http.domain"),
					Port:    viper.GetInt("http.port"),
					Root:    pwd,
					Version: "v1",
				})
			}),
		},
		{
			Name:    "cache",
			Aliases: []string{"c"},
			Usage:   "cache operations",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Usage:   "list all cache keys",
					Aliases: []string{"l"},
					Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
						keys, err := p.Cache.Keys()
						if err != nil {
							return err
						}
						for _, k := range keys {
							fmt.Println(k)
						}
						return nil
					}),
				},
				{
					Name:    "clear",
					Usage:   "clear cache items",
					Aliases: []string{"c"},
					Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
						return p.Cache.Flush()
					}),
				},
			},
		},
		{
			Name:    "oauth",
			Aliases: []string{"oa"},
			Usage:   "oauth credentials.",
			Subcommands: []cli.Command{
				{
					Name:    "google",
					Usage:   "google oauth v2",
					Aliases: []string{"g"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "file, f",
							Value: "",
							Usage: "google oauth2 json filename.",
						},
					},
					Action: web.IocAction(func(c *cli.Context, g *inject.Graph) error {
						fn := c.String("file")
						if fn == "" {
							return errors.New("filename mustn't empty")
						}

						fd, err := os.Open(fn)
						if err != nil {
							return err
						}
						defer fd.Close()

						var gc GoogleCredential
						dec := json.NewDecoder(fd)
						if err := dec.Decode(&gc); err != nil {
							return err
						}
						return p.Dao.Set("google.oauth", gc.To(), true)
					}),
				},
			},
		},
		{
			Name:    "users",
			Aliases: []string{"us"},
			Usage:   "users manager",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Usage:   "list all users",
					Aliases: []string{"l"},
					Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
						var us []User
						if err := p.Dao.Db.
							Select([]string{"uid", "name", "email"}).
							Order("last_sign_in DESC").
							Find(&us).Error; err != nil {
							return err
						}
						fmt.Println("UID\tUSER")
						for _, u := range us {
							fmt.Printf("%s\t%s<%s>\n", u.UID, u.Name, u.Email)
						}
						return nil
					}),
				},
				{
					Name:    "role",
					Usage:   "add/remove user's role",
					Aliases: []string{"r"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "uid, u",
							Value: "",
							Usage: "user's uid",
						},
						cli.StringFlag{
							Name:  "name, n",
							Value: "",
							Usage: "role's name",
						},
						cli.BoolFlag{
							Name:  "deny, d",
							Usage: "remove role from user",
						},
						cli.IntFlag{
							Name:  "years, y",
							Value: 10,
							Usage: "years",
						},
					},
					Action: web.IocAction(func(c *cli.Context, _ *inject.Graph) error {
						uid := c.String("uid")
						name := c.String("name")
						deny := c.Bool("deny")
						years := c.Int("years")
						if uid == "" {
							return errors.New("uid mustn't empty.")
						}
						if name == "" {
							return errors.New("role's name mustn't empty.")
						}
						user, err := p.Dao.GetUser(uid)
						if err != nil {
							return err
						}
						role, err := p.Dao.Role(name, "-", 0)
						if err != nil {
							return err
						}
						if deny {
							err = p.Dao.Deny(role.ID, user.ID)
						} else {
							err = p.Dao.Allow(role.ID, user.ID, years, 0, 0)
						}
						return err
					}),
				},
			},
		},
	}
}

func init() {
	viper.SetDefault("workers.queues.email", 5)
}
