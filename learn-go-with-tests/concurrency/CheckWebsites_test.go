package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waa://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com/",
		"http://bolg.gypsydave5.com/",
		"waa://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com/":          true,
		"http://bolg.gypsydave5.com/": true,
		"waa://furhurterwe.geds":      false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v got %v", want, got)
	}
}
