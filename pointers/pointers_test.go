package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Test Bitcoin deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		got := wallet.Balance()
		want := Bitcoin(10.0) //Cast the float amount to Bitcoin type

		if got != want {
			t.Errorf("got %s want %s. %v", got, want, wallet)
		}
	})

	t.Run("Test Withdraw Bitcoin", func(t *testing.T) {
		wallet := Wallet{15.0}
		wallet.Withdraw(Bitcoin(10.0))
		got := wallet.Balance()
		want := Bitcoin(5.0)

		if got != want {
			t.Errorf("got %s want %s. %v", got, want, wallet)
		}
	})

	t.Run("Test printing balance", func(t *testing.T) {
		wallet := Wallet{15.0}
		got := wallet.Balance().String()
		want := "15.00 BTC"

		if got != want {
			t.Errorf("got %s want %s. %v", got, want, wallet)
		}
	})
}
