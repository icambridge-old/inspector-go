package controllers

import (
	"database/sql"
	_ "code.google.com/p/go-mysql-driver/mysql"
	"fmt"
	"github.com/robfig/revel"
)

var (
	db *sql.DB
)

type DbPlugin struct {
	rev.EmptyPlugin
}

func (p DbPlugin) OnAppStart() {
	
	
}

func (p DbPlugin) BeforeRequest(c *rev.Controller) {
	fmt.Println("Hello")
	var err error

	dsn, found := rev.Config.String("mysql.dsn")

	if found != true {
		panic("MySQL dsn not found")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}	
	txn, err := db.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
}

func (p DbPlugin) AfterRequest(c *rev.Controller) {
	if err := c.Txn.Commit(); err != nil {
		if err != sql.ErrTxDone {
			panic(err)
		}
	}
	c.Txn = nil
}

func (p DbPlugin) OnException(c *rev.Controller, err interface{}) {
	if c.Txn == nil {
		return
	}
	if err := c.Txn.Rollback(); err != nil {
		if err != sql.ErrTxDone {
			panic(err)
		}
	}
}

func init() {
	rev.RegisterPlugin(DbPlugin{})
}
