package iteration

func Repeat(s string, repeatCount int) string {
  var result string
  for range(repeatCount) {
    result += s
  }
  return result
}
