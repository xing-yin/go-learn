package point_error

import (
	"errors"
	"fmt"
)

// Bitcoin 自定义类型：比 int 更有描述性
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// Deposit 订金存入
func (w *Wallet) Deposit(amount Bitcoin) {
	// 当调用一个函数或方法时，参数会被复制。(参数为 w Wallet) ===> 需要用指针指向同一个地址（balance）
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

// Balance 账户余额
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// InsufficientFundsErrors ==> 将错误重构为一个变量，并为其提供一个单一的事实来源
// 重构：我们并不关心具体的措辞是什么，只是在给定条件的情况下返回一些有意义的错误
var InsufficientFundsErrors = errors.New("cannot withdraw, insufficient funds")

// Withdraw 提取
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return InsufficientFundsErrors
	}
	w.balance -= amount
	return nil
}

// 实现 Bitcoin 的 Stringer 方法（位于 fmt 包中定义的）==> 当使用 %s 打印格式化的字符串时，可以定义此类型的打印方式
// 实现了 fmt.Stringer
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
