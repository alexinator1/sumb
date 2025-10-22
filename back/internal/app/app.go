package app

import (
	"fmt"
)

type App struct {
	Cfg      *AppConfig
	modulses *[]Module
}

func NewApp(cfg *AppConfig, modules ...Module) *App {
	return &App{
		Cfg:      cfg,
		modulses: modules,
	}
}

func (a *App) addModule(module Module) {
	*a.modulses = append(*a.modulses, module)
}

func Run() {
	fmt.Println("Hello, World!")
	cfg := Load()
	fmt.Println(cfg)
}
