package signal

import (
	"os"
	"os/signal"

)

// WaitForSignal waits until one of the provided signals are occurring.
// The function returns the signal that was received.
func WaitForSignal(signals ...os.Signal) os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, signals...)
	return <-sigChan
}
