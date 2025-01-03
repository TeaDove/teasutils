package notify_utils

import (
	"os"
	"os/signal"
)

func RunOnInterrupt(f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			f()
		}
	}()
}
