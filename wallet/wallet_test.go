package wallet

import "testing"

func TestWallet(t *testing.T) {

  assertBalance := func (t testing.TB, w Wallet, want Bitcoin)  {
    if w.balance != want {
			t.Errorf("got %s want %s", w.balance, want)
    }
  }

  assertError := func (t testing.TB, err, want error)  {
    t.Helper()
    if err == nil {
      t.Error("wanted an error but didn't get one")
    }
    if err != want {
      t.Errorf("got %q want %q", err, want)
    }
  }

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)

    assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
    err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

    assertBalance(t, wallet, want)
    if err != nil {
      t.Errorf("want no errors, but got %q", err)
    }
	})

	t.Run("withdraw insufficient balance", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		want := Bitcoin(20)
    err := wallet.Withdraw(Bitcoin(100))

    assertBalance(t, wallet, want)
    assertError(t, err, ErrInsuffientFunds)
	})
}
