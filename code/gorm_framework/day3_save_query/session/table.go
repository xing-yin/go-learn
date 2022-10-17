package session

import (
	"fmt"
	"geeorm/log"
	"geeorm/schema"
	"reflect"
	"strings"
)

// Model 存放操作数据库表相关的代码
// 用于给 refTable 赋值。解析操作是比较耗时的，因此将解析的结果保存在成员变量 refTable 中，即使 Model() 被调用多次，如果传入的结构体名称不发生变化，则不会更新 refTable 的值。
func (s *Session) Model(value interface{}) *Session {
	// nil or different model,update refTable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// RefTable 返回 refTable 的值，如果 refTable 未被赋值，则打印错误日志
func (s Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

// CreateTable 三个方法的实现逻辑是相似的，利用 RefTable() 返回的数据库表和字段的信息，拼接出 SQL 语句，调用原生 SQL 接口执行
func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Type))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row, _ := s.Raw(sql, values...).QueryRows()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
