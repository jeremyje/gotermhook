// +build appengine

package gotermhook

import (
	"os"
	"os/signal"
)

func addTerminatingSignals(c chan<- os.Signal) {
	signal.Notify(c, os.Interrupt)
}
