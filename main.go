package main

import "github.com/ipoluianov/tfm/goforms"

func main() {
	a := goforms.NewApp()
	f := goforms.NewForm()
	a.ExecMainForm(f)
}
