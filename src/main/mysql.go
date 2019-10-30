package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"models"
)

var engine *xorm.Engine

func init() {
	var err error
	dataSourceName := "root:TLRGD0aF1FtyEd25@tcp(172.17.20.67:3307)/news_data"
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatal(fmt.Sprintf("connect mysql fail ! [%s]", err))
	}
}

func InsertNews(news [] models.TmpNewsCrawalContent) (int64, bool) {
	engine.TableName("asd")
	affected, err := engine.Insert(news)
	if err != nil {
		fmt.Println(err)
		return affected, false
	}
	return affected, true
}
