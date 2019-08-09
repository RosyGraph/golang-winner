package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {

	slowServer := makeDelayedServer(time.Millisecond(20))
	fastServer := makeDelayedServer(time.Millisecond(10))

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
