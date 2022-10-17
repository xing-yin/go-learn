package session

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/schema"
	"strings"
)

type Session struct {
	// 即使用 sql.Open() 方法连接数据库成功之后返回的指针
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	// 用来拼接 SQL 语句和
	sql strings.Builder
	// 用来拼接 SQL 语句中占位符的对应值
	sqlVars []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw appends sql and sqlVars
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// 封装 Exec()、Query() 和 QueryRow() 三个原生方法
// 封装有 2 个目的:一是统一打印日志（包括执行的 SQL 语句和错误日志）; 二是执行完成后，清空 (s *Session).sql 和 (s *Session).sqlVars 两个变量。这样 Session 可以复用，开启一次会话，可以执行多次 SQL。

// Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil { // ignore_security_alert
		log.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...) // ignore_security_alert
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	rows, err = s.DB().Query(s.sql.String(), s.sqlVars...) // ignore_security_alert
	if err != nil {
		log.Error(err)
	}
	return
}
