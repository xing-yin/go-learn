package employee

import "testing"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/9
 * @Desc:
 */

//-------------------------------------------
type fakeStmtForMaleCount struct {
	Stmt
}

func (fakeStmtForMaleCount) Exec(stmt string, args ...string) (Result, error) {
	return Result{Count: 5}, nil
}

func TestEmployeeMaleCount(t *testing.T) {
	f := fakeStmtForMaleCount{}
	c, _ := MaleCount(f)
	if c != 5 {
		t.Errorf("want: %d,actual: %d", 5, c)
		return
	}
}
