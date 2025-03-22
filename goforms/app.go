package goforms

import "fmt"

type App struct {
}

func NewApp() *App {
	var c App
	return &c
}

func (c *App) ExecMainForm(form *Form) {
	fmt.Println("App.ExecMainForm()")
	form.Show()

	Run()
}
