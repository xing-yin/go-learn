package clause

import (
	"reflect"
	"testing"
)

func TestClause_Build(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		testSelect(t)
	})
}

func testSelect(t *testing.T) {
	var caluse Clause
	caluse.Set(LIMIT, 3)
	caluse.Set(SELECT, "User", []string{"*"})
	caluse.Set(WHERE, "Name = ?", "Bob")
	caluse.Set(ORDERBY, "Age ASC")
	sql, vars := caluse.Build(SELECT, WHERE, ORDERBY, LIMIT)
	t.Log(sql, vars)
	if sql != "SELECT * FROM User WHERE Name = ? ORDER BY Age ASC LIMIT ?" {
		t.Fatal("failed to build SQL")
	}
	if !reflect.DeepEqual(vars, []interface{}{"Bob", 3}) {
		t.Fatal("failed to build SQLVars")
	}
}
