package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Wanted an error but didn't get one")
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		/*		got := wallet.Balance()
				want := Bitcoin(10)

				if got != want {
					t.Errorf("got %s want %s", got, want)
				}*/

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		/*got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}*/

	})
	t.Run("withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
		/*if err == nil {
			t.Error("Wnated an error but didn't get one")
		}*/
	})
}
