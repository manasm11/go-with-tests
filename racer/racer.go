package racer

import "net/http"

func Racer(urlA, urlB string) (winner string) {
  select {
  case <- ping(urlA):
  return urlA
  case <- ping(urlB):
  return urlB
  }
//  aDuration := measureResponseTime(urlA)
	// bDuration := measureResponseTime(urlB)

	// if aDuration < bDuration {
		// return urlA
	// }
	// return urlB
}

func ping(url string) chan struct{} {
  ch := make(chan struct{})
  go func(u string) {
    http.Get(u)
    close(ch)
  }(url)
  return ch
}
// func measureResponseTime(url string) time.Duration {
	// start := time.Now()
	// http.Get(url)
	// return time.Since(start)
// }
