package racer

import (
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

type RacerError string

func (r RacerError) Error() string {
	return string(r)
}

const (
	RacerTimeExceeded = RacerError("response time exceeded 10 seconds")
)

func Racer(urlA, urlB string) (winner string, err error) {
  return ConfigurableRacer(urlA, urlB, tenSecondTimeout)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", RacerTimeExceeded
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func(u string) {
		http.Get(u)
		close(ch)
	}(url)
	return ch
}
