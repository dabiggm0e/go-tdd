package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet *Wallet, want Bitcoin) {
		t.Helper()
		if wallet.Balance() != want {
			t.Errorf("got %s want %s. %v", wallet.Balance(), want, wallet)
		}
	}

	t.Run("Test Bitcoin deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		want := Bitcoin(10.0) //Cast the float amount to Bitcoin type

		assertBalance(t, &wallet, want)
	})

	t.Run("Test Withdraw Bitcoin", func(t *testing.T) {
		wallet := Wallet{15.0}
		wallet.Withdraw(Bitcoin(10.0))
		want := Bitcoin(5.0)

		assertBalance(t, &wallet, want)
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
