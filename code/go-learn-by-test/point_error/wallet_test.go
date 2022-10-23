package point_error

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(15))
		assertNotError(t, err)
	})

	// 模拟超额提取
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, InsufficientFundsErrors)
	})
}

func assertNotError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but did not want one")
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("want an error but didn't get one")
	}

	if got != want {
		t.Errorf("got '%s',want '%s'", got, want)
	}
}
