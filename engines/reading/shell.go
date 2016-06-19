package reading

import (
	"github.com/facebookgo/inject"
	"github.com/itpkg/chaos/web"
	"github.com/urfave/cli"
)

//Shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{

		{
			Name:    "reading",
			Aliases: []string{"rd"},
			Usage:   "reading engine operations",
			Subcommands: []cli.Command{
				{
					Name:    "scan",
					Usage:   "scan books",
					Aliases: []string{"s"},
					Action: web.IocAction(func(*cli.Context, *inject.Graph) error {
						return p._scanBooks()
					}),
				},
			},
		},
	}
}
