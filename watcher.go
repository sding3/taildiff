package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type watcher struct {
	interval  time.Duration
	command   string
	output    ChangeDetectingBuffer
	exitOnErr bool
}

func (w *watcher) Start() int {
	tick := setupTicker(w.interval)

	for {
		pState := w.execute()
		if w.exitOnErr && !pState.Success() {
			return pState.ExitCode()
		}

		<-tick
	}
}

func (w *watcher) execute() *os.ProcessState {
	cmd := exec.Command(getShell(), []string{"-c", w.command}...)
	cmd.Stdout = &w.output
	cmd.Run() // err discarded
	w.output.Done()

	if w.output.Changed {
		fmt.Println(&w.output)
	}

	w.output.Rewind()
	return cmd.ProcessState
}

func setupTicker(d time.Duration) <-chan time.Time {
	if d <= 0 {
		c := make(chan time.Time)
		close(c)
		return c
	}
	return time.NewTicker(d).C
}

func getShell() string {
	envShell := os.Getenv("SHELL")

	if envShell == "" {
		return "/bin/sh"
	}
	return envShell
}
