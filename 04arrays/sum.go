package sums
func Sum(nums []int) int {
  var result int
  for _, n := range nums {
    result += n
  }
  return result
}

func SumAll(a, b []int) []int {
  return []int{Sum(a), Sum(b)}
}
