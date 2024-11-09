package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
  repeated := Repeat("a", 9)
  expected := "aaaaaaaaa"

  if repeated != expected {
    t.Errorf("expected %q but got %q", expected, repeated)
  }
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a", 6)
  }
}

func ExampleRepeat() {
  fmt.Println(Repeat("ab", 3))
  //Output:ababab
}
