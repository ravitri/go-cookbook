package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func CatchSig(ch chan os.Signal, done chan bool) {

	sig := <-ch
	fmt.Println("\nSignal Received:", sig)

	switch sig {
	case syscall.SIGINT:
		fmt.Println("handlig SIGINT signal now")
	case syscall.SIGTERM:
		fmt.Println("handling SIGTERM signal now")
	default:
		fmt.Println("Unexpected signal received")
	}

	done <- true
}

func main() {

	// initialize the channels
	signals := make(chan os.Signal, 2)
	done := make(chan bool)

	// Relay the SIGINT and SIGTERM calls to channel if caught
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// If a signal is caught by this go routine, it will write to done
	go CatchSig(signals, done)

	fmt.Println("Press Ctrl-c to terminate...")
	<-done
	fmt.Println("Done!")
}
