package auth

import "github.com/urfave/cli"

var ENV = cli.StringFlag{
	Name:   "lang, e",
	Value:  "english",
	Usage:  "language for the greeting",
	EnvVar: "LEGACY_COMPAT_LANG,APP_LANG,LANG",
}
