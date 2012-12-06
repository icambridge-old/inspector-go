package controllers

import "github.com/robfig/revel"

type Application struct {
	*rev.Controller
}

func (c Application) Index() rev.Result {
	random := "Test"
	decid := "Second"
	return c.Render(random, decid)
}

