package web

type Engine interface {
	Mount()
	Migrate()
	Seed()
	Worker()
	Shell()
}

var engines []Engine

func Register(ens ...Engine) {
	engines = append(engines, ens...)
}
