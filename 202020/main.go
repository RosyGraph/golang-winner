package main

import (
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"time"
)

func main() {
	breakNote := gosxnotifier.NewNotification("Stare away for 20 seconds")
	breakNote.Sound = gosxnotifier.Basso
	resumeNote := gosxnotifier.NewNotification("Resume work")
	resumeNote.Sound = gosxnotifier.Basso
	for {
		fmt.Printf("Starting program at %v. Ctrl-C to exit.\n", time.Now())
		time.Sleep(20 * 60 * time.Second)
		breakNote.Push()
		time.Sleep(20 * time.Second)
		resumeNote.Push()
	}
}
