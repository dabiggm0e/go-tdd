package pointers

import "testing"

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10.0))
	got := wallet.Balance()
	want := Bitcoin(10.0) //Cast the float amount to Bitcoin type

	if got != want {
		t.Errorf("want %.2f got %.2f. %v", got, want, wallet)
	}
}
