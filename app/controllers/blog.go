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

	rows, err := c.Txn.Query("SELECT `title`, `description`,`slug`,`timestamp` FROM `posts` ORDER BY id DESC")
	if err != nil {
		panic(err)
	}

	posts := []map[string] string{}

	var title string
	var description string
	var slug string
	var post map[string] string
	var date string

	for rows.Next() {
 		err = rows.Scan( &title, &description, &slug, &date)
		post = map[string] string {
			"Title": title,
			"Description": description,
			"Slug": slug,
			"Date": date,
		}	
		posts = append(posts, post)
	}

	return c.Render(posts)
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
