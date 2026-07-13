package notifyutils

import (
	"os"
	"os/signal"
)

// OnInterrupt runs f the first time an os.Interrupt (Ctrl+C) is received, then
// restores the default signal handling so a second interrupt terminates the
// process. It returns immediately; f runs on a background goroutine.
func OnInterrupt(f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for sig := range c {
			f()
			signal.Reset(sig)
		}
	}()
}
