package main

import (
	"runtime"

	"github.com/ipoluianov/tfm/goforms"
)

func main() {
	runtime.LockOSThread()
	a := goforms.NewApp()
	f := goforms.NewForm()
	a.ExecMainForm(f)
}
