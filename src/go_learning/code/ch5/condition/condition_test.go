package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	// if 支持变量量赋值
	if a := 1 == 1; a {
		// 主要是为了支持多返回值赋值
		//if v,err := someFun(); err {
		t.Log("1==1")
	}
}

// switch条件：
// a.条件表达式不限制为常量或者整数;
// b.单个 case 中，可以出现多个结果选项, 使⽤用逗号分隔;
// c.与 C 语⾔等规则相反，Go 语⾔不需要用break来明确退出⼀一个 case;
// d.可以不设定 switch 之后的条件表达式，在此种情况下，整个 switch 结
// e.构与多个 if...else... 的逻辑作⽤用等同
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		// b
		case 0, 2:
			t.Log("Even")
			//	c
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

// 简化 if else
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		// 没有条件的switch,相当于简化的 if else
		switch {
		// case 后面直接跟条件
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unKnow")
		}
	}
}
