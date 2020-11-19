package pointers_errors

import (
	"testing"
)

func TestBitcoin(t *testing.T) {
	amount := Bitcoin(10)
	got := amount.String()
	expected := "10 BTC"

	if got != expected {
		t.Errorf("got %s expected %s", got, expected)
	}
}

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		got := wallet.Balance()

		if got != expected {
			t.Errorf("got %s expected %s", got, expected)
		}
	}

	assertError := func(t *testing.T, err error, expected error) {
		t.Helper()

		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}

		if err != expected {
			t.Errorf("expected %q got %q", expected, err)
		}
	}

	assertNotError := func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Fatal("got an error but didn't expected one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNotError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, initialBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
