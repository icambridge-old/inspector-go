package controllers

import (
	"github.com/robfig/revel"
)
type Error struct {
	*rev.Controller
}

func (c Error) Index() rev.Result {
	return c.Render()
}
