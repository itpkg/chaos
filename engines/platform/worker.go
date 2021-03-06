package platform

import (
	"github.com/jrallison/go-workers"
	"github.com/spf13/viper"
)

//Worker worker job
func (p *Engine) Worker() {
	workers.Process("email", func(msg *workers.Msg) {

		p.Logger.Infof("GET JOB %s@email", msg.Jid())
		// args, err := msg.Array()
		// if err != nil {
		// 	p.Logger.Error(err)
		// }
		p.Logger.Debugf("ARGS: %+v", msg.Interface().(map[string]interface{}))

	}, viper.GetInt("workers.queues.email"))
}
