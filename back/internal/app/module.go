package app

type Module interface {
	Name() string
	Init(app *App) error
	New(app *App) Module
}
