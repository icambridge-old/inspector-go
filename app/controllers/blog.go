package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	"html/template"
)
type Blog struct {
	*rev.Controller
}

func (c Blog) Index() rev.Result {
	random := "Test"
	decid := "Second"
	return c.Render(random, decid)
}

func (c Blog) View() rev.Result {

	var permalink string = c.Params.Get("permalink")

	st, err := c.Txn.Prepare("SELECT `title`,`content` FROM `posts` WHERE slug = ?")
	if err != nil {
		fmt.Println(err)
		return c.Render(permalink)
	}


	rows, err := st.Query(permalink)
	if err != nil {
		fmt.Println(err)
		return c.Render(permalink)
	}


	var title string
	var rawContent string
	rows.Next()
	err = rows.Scan( &title, &rawContent)
	content := template.HTML(rawContent)

	return c.Render(permalink, title, content)
}
