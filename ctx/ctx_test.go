package ctx

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		var (
			data  = "hello, world"
			store = &SpyStore{response: data}
			svr   = Server(store)
		)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s" want "%s"`, response.Body.String(), data)
		}

		if store.cancelled {
      t.Error("store should not have been cancelled")
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		var (
			data     = "hello, world"
			store    = &SpyStore{response: data}
			svr      = Server(store)
			request  = httptest.NewRequest(http.MethodGet, "/", nil)
			response = httptest.NewRecorder()
		)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("store was not cancelled")
		}
	})
}
