package employee

import "back/internal/app"

type EmployeeModule struct {
	name string
}

func NewEmployeeModule() *EmployeeModule {
	return &EmployeeModule{
		name: "EmployeeModule",
	}
}

func (m *EmployeeModule) Name() string {
	return m.name
}

func (m *EmployeeModule) Init(a *app.App) error {
	// Initialization logic for the Employee module
	return nil
}
