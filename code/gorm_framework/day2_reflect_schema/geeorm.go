package geeorm

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/session"
)

// Session 负责与数据库的交互，那交互前的准备工作（比如连接/测试数据库），交互后的收尾工作（关闭连接）等就交给 Engine 来负责了。
// Engine 是 GeeORM 与用户交互的入口
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

// NewEngine 连接数据库，返回 *sql.DB
//NewEngine 创建 Engine 实例时，获取 driver 对应的 dialect
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	// make sure the specific dialect exists
	dialect, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}

	e = &Engine{db: db, dialect: dialect}
	log.Info("Connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

// NewSession 通过 Engine 实例创建会话，进而与数据库进行交互了(创建 Session 实例时，传递 dialect 给构造函数 New)
func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
