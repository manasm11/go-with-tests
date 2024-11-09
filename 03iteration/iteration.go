package iteration

const repeatCount = 5

func Repeat(s string) string {
  var result string
  for range(repeatCount) {
    result += s
  }
  return result
}
