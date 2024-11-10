package sums
func Sum(nums []int) int {
  var result int
  for _, n := range nums {
    result += n
  }
  return result
}

func SumAll(arrs...  []int) []int {
  result := []int{}
  for _, arr := range arrs {
    result = append(result, Sum(arr))
  }
  return result
}
