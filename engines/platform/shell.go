package platform

import (
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
			Action: func(*cli.Context) error {
				const tpl = `

server {
  listen {{if .Ssl}}443{{- else}}80{{- end}};
{{if .Ssl}}
  ssl  on;
  ssl_certificate  ssl/www.{{.Domain}}-cert.pem;
  ssl_certificate_key  ssl/www.{{.Domain}}-key.pem;
  ssl_session_timeout  5m;
  ssl_protocols  SSLv2 SSLv3 TLSv1;
  ssl_ciphers  RC4:HIGH:!aNULL:!MD5;
  ssl_prefer_server_ciphers  on;
{{- end}}
  client_max_body_size 4G;
  keepalive_timeout 10;

  server_name www.{{.Domain}};

  root {{.Root}}/public;
  index index.html;

  access_log log/www.{{.Domain}}.access.log;
  error_log log/www.{{.Domain}}.error.log;

  location / {
    try_files $uri $uri/ /index.html;
  }

  location ^~ /assets/ {
    gzip_static on;
    expires max;
    access_log off;
    add_header Cache-Control public;
  }
}

server {
  listen {{if .Ssl}}443{{- else}}80{{- end}};
{{if .Ssl}}
  ssl  on;
  ssl_certificate  ssl/api.{{.Domain}}-cert.pem;
  ssl_certificate_key  ssl/api.{{.Domain}}-key.pem;
  ssl_session_timeout  5m;
  ssl_protocols  SSLv2 SSLv3 TLSv1;
  ssl_ciphers  RC4:HIGH:!aNULL:!MD5;
  ssl_prefer_server_ciphers  on;
{{- end}}
  client_max_body_size 4G;
  keepalive_timeout 10;

  server_name api.{{.Domain}};
  access_log log/api.{{.Domain}}.access.log;
  error_log log/api.{{.Domain}}.error.log;

  location / {
    {{if .Ssl}}proxy_set_header  X-Forwarded-Proto https;{{- end}}
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_redirect off;
    proxy_pass http://localhost:{{.Port}};
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
					Domain string
					Port   int
					Ssl    bool
					Root   string
				}{
					Ssl:    viper.GetBool("http.ssl"),
					Domain: viper.GetString("http.domain"),
					Port:   viper.GetInt("http.port"),
					Root:   pwd,
				})
			},
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
	}
}

func init() {
	viper.SetDefault("workers.queues.email", 5)
}
