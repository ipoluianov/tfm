package goforms

import "fmt"

type Form struct {
}

func NewForm() *Form {
	var c Form
	return &c
}

func (c *Form) Show() {
	fmt.Println("Form.Show()")
}
