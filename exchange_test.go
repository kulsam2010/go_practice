package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("testing deposit ", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(t, got, want)
	})

	t.Run("testing withdrawl ", func(t *testing.T) {
		wallet := Wallet{Bitcoin(200)}
		err := wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(190)
		assertBalance(t, got, want)
		assetNoError(t, err)
	})

	t.Run("testing withdrawl insufficient balance ", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		startingBalance := wallet.Balance()
		err := wallet.Withdraw(Bitcoin(20))
		got := wallet.Balance()
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, got, startingBalance)
	})
}

func assertBalance(t testing.TB, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func assetNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("wanted nil error but got %q", err)
	}
}
