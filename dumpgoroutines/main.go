// package dumpgoroutines provides a simple way to have your program print
// current goroutine stack traces upon receiving a signal
package dumpgoroutines

import (
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

// Installs with default signal (SIGUSR2)
func InstallDefault() {
	Install(syscall.SIGUSR2)
}

// Installs with provided signals (see signal.Notify)
func Install(signals ...os.Signal) {
	c := make(chan os.Signal)
	signal.Notify(c, signals...)
	go func() {
		for {
			<-c
			pprof.Lookup("goroutine").WriteTo(os.Stderr, 1)
		}
	}()
}
