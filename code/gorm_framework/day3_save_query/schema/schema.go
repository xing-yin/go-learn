package schema

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
)

// Field represents a column of database
type Field struct {
	// 字段名
	Name string
	// 类型
	Type string
	// 约束条件
	Tag string
}

// Schema represents a table of database
type Schema struct {
	// 被映射的对象
	Model interface{}
	// 表名
	Name string
	// 字段
	Fields []*Field
	// 包含所有的字段名（列名）
	FieldNames []string
	// 记录字段名和 Filed 的映射关系，方便直接使用
	fieldMap map[string]*Field
}

func (s *Schema) GetFiled(name string) *Field {
	return s.fieldMap[name]
}

// Parse 将任意的对象解析为 Schema 实例
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// TypeOf() 和 ValueOf() 是 reflect 包最为基本也是最重要的 2 个方法，分别用来返回入参的类型和值。
	//因为设计的入参是一个对象的指针，因此需要 reflect.Indirect() 获取指针指向的实例。modelType.Name() 获取到结构体的名称作为表名。
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				// 通过 (Dialect).DataTypeOf() 转换为数据库的字段类型
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

// RecordValues 将 user1、user2 转换为 ("Tom", 18), ("Same", 25) 这样的格式
func (s *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range s.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}
