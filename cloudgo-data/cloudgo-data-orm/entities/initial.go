package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var mydb *xorm.Engine

// connect the database
func init() {
	db, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	err = db.Sync2(new(UserInfo))
	checkErr(err)
	mydb = db
}

// SQLExecer interface for supporting xorm.Engine and xorm.Session to do sql statement
type SQLExecer interface {
	Find(beans interface{}, condiBeans ...interface{}) error
	Insert(beans ...interface{}) (int64, error)
	Where(query interface{}, args ...interface{}) *xorm.Session
}

// DaoSource Data Access Object Source
type DaoSource struct {
	SQLExecer
}

// error detection
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
