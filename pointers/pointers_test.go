package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Test Bitcoin deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		want := Bitcoin(10.0) //Cast the float amount to Bitcoin type

		assertBalance(t, &wallet, want)
	})

	t.Run("Test Withdraw Bitcoin", func(t *testing.T) {
		wallet := Wallet{15.0}
		err := wallet.Withdraw(Bitcoin(10.0))
		want := Bitcoin(5.0)
		assertNoError(t, &wallet, err)
		assertBalance(t, &wallet, want)
	})

	t.Run("Test overdraft Withdraw Bitcoin", func(t *testing.T) {
		startingBalance := Bitcoin(15.0)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(20.0))

		assertBalance(t, &wallet, startingBalance)
		assertError(t, &wallet, err, ErrInsufficientBalance)

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

func assertBalance(t *testing.T, wallet *Wallet, want Bitcoin) {
	t.Helper()
	if wallet.Balance() != want {
		t.Errorf("got %s want %s. %v", wallet.Balance(), want, wallet)
	}
}

func assertError(t *testing.T, wallet *Wallet, got error, want error) {
	if got == nil {
		t.Fatalf("Didn't get any error! Need '%v'. %v", want, wallet)
	}

	if got != want {
		t.Errorf("got err='%v' need '%v'. %v", got, want, wallet)
	}
}

func assertNoError(t *testing.T, wallet *Wallet, got error) {

	if got != nil {
		t.Errorf("got err='%v' but didn't need one. %v", got, wallet)
	}
}
