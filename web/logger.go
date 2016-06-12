package web

import (
	"os"

	"github.com/op/go-logging"
)

//Logger open logger
func Logger() *logging.Logger {
	var bkd logging.Backend
	if IsProduction() {
		var err error
		bkd, err = logging.NewSyslogBackend("itpkg")
		if err != nil {
			bkd = logging.NewLogBackend(os.Stdout, "", 0)
		}
	} else {
		bkd = logging.NewLogBackend(os.Stdout, "", 0)
	}

	//`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`
	if IsProduction() {
		logging.SetFormatter(logging.MustStringFormatter(`%{color}%{level:.4s} %{id:03x} %{color:reset} [%{shortfunc}] %{message}`))
		logging.SetLevel(logging.INFO, "")
	} else {
		logging.SetFormatter(logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{level:.4s} %{id:03x} %{color:reset} [%{longfunc}] %{message}`))
	}
	logging.SetBackend(bkd)
	return logging.MustGetLogger("chaos")
}
