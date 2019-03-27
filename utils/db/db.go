package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

//定义单例
var G_db *xorm.Engine

type Db struct {
	
}

/**
 实例化数据库
 */
func New(dbtype string,dbDns string ) (err error) {
	connection, err := xorm.NewEngine(dbtype,dbDns)
	if err != nil {
		panic("数据库连接错误")
	}
	G_db = connection
	return err;
}