package sums
func Sum(nums []int) int {
  var result int
  for _, n := range nums {
    result += n
  }
  return result
}
