package ops

import "github.com/urfave/cli"

//FlagDatabaseUser database user flag
var FlagDatabaseUser = cli.StringFlag{
	Name:  "user, u",
	Value: "",
	Usage: "database user",
}

//FlagDatabasePassword database password flag
var FlagDatabasePassword = cli.StringFlag{
	Name:  "password, p",
	Value: "",
	Usage: "database password",
}
