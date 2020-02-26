package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	w, err := NewWatcherFromFlag()
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	os.Exit(w.Start())
}

func NewWatcherFromFlag() (*watcher, error) {
	w := &watcher{}

	flag.StringVar(&w.command, "c", "", "[required] shell command")
	flag.DurationVar(&w.interval, "n", time.Second, "update interval")
	flag.BoolVar(&w.exitOnErr, "e", false, "exit on command error (default false)")
	flag.Parse()
	if w.command == "" {
		return nil, fmt.Errorf("A command is required. Supply command with -c <command>")
	}

	return w, nil
}
