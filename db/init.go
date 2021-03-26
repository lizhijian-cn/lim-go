package db

import (
	"github.com/aceld/zinx/zlog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"xorm.io/core"
)

var (
	Engine *xorm.Engine
)

func InitDB() {
	engine, err := xorm.NewEngine("mysql", "root:lizhijian1098@/LIM?charset=utf8")
	if err != nil {
		zlog.Fatalf("failed to connect mysql: %s", err)
		return
	}
	//f, err := os.Open("log/sql.log")
	//defer f.Close()
	if err != nil {
		zlog.Fatalf("failed to create sql.log")
		return
	}
	Engine = engine
	logger := xorm.NewSimpleLogger(os.Stdout)
	logger.ShowSQL(true)
	logger.SetLevel(core.LOG_INFO)
	Engine.SetLogger(logger)

	if err := engine.Ping(); err != nil {
		zlog.Fatalf("failed to ping database: %s", err)
		return
	}
}
