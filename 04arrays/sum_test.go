package sums

import "testing"

func TestSum(t *testing.T) {

	t.Run("sum collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 4}

		got := Sum(numbers)
		want := 7

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
  // helper function
	sliceEqual := func(t testing.TB, a, b []int) bool {
		t.Helper()
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

  // test summing two slices
	t.Run("sum two slices", func(t *testing.T) {
		got := SumAll([]int{1, 3}, []int{0, 9, 1})
		want := []int{4, 10}

		if !sliceEqual(t, got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

  // test summing three slices
	t.Run("sum three slices", func(t *testing.T) {
		got := SumAll([]int{1, 3}, []int{0, 9}, []int{5, 8})
		want := []int{4, 9, 13}

		if !sliceEqual(t, got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
