package utils

import (
	"os"
	"os/signal"
)

func WaitSignals(job func(), signals ...os.Signal) {
	c := make(chan os.Signal)
	signal.Notify(c, signals...)
	go job()
	<-c
}
