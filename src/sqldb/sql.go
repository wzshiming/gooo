package sqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

type SqlDB struct {
	sql.DB
	Tag string
}

func NewSqlDBFrom(source string) *SqlDB {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err.Error())
	}
	return &SqlDB{
		DB:  *db,
		Tag: "dbname",
	}
}

func NewSqlDB(name, pwd, addr, dbname string) *SqlDB {
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", name, pwd, addr, dbname)
	return NewSqlDBFrom(source)
}

func (s *SqlDB) InsertPrepare(tables string, key []string) (*sql.Stmt, error) {
	var t string
	for _, v := range key {
		t = fmt.Sprintf("%s,%s=?", t, v)
	}
	querys := fmt.Sprintf("insert %s set %s", tables, t[1:])
	return s.Prepare(querys)
}

func (s *SqlDB) Insert(i interface{}) (id int64, err error) {
	v := reflect.ValueOf(i)
	t := v.Type()
	var st string
	var is []interface{}
	for i := 0; i != t.FieldAlign(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get(s.Tag)
		if tag != "" {
			st = fmt.Sprintf("%s,%s=?", st, tag)
			is = append(is, v.Field(i).Interface())
		}
	}
	querys := fmt.Sprintf("insert %s set %s", t.Name(), st[1:])
	res, err := s.Exec(querys, is...)
	if err != nil {
		return
	}
	return res.LastInsertId()
}

type SqlTable struct {
	sql.Stmt
}
