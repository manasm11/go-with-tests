package wallet

import "testing"

func TestWallet(t *testing.T) {

  assertBalance := func (t testing.TB, w Wallet, want Bitcoin)  {
    if w.balance != want {
			t.Errorf("got %s want %s", w.balance, want)
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
    wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

    assertBalance(t, wallet, want)
	})
}
