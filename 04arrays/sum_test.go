package sums

import "testing"

func TestSum(t *testing.T) {

  t.Run("sum collection of any size", func(t *testing.T) {
    numbers := []int{1, 2, 4}

    got :=Sum(numbers)
    want := 7

    if got != want {
      t.Errorf("got %d want %d given %v", got, want, numbers)
    }
  })
}
