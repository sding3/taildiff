package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	versionMajor = 0
	versionMinor = 0
	versionPatch = 1
)

var CommandWatcher = &watcher{}

func init() {
	flag.ErrHelp = fmt.Errorf("flag error")
	flag.StringVar(&CommandWatcher.command, "c", "", "[required] shell command")
	flag.DurationVar(&CommandWatcher.interval, "n", time.Second, "update interval")
	flag.BoolVar(&CommandWatcher.exitOnErr, "e", false, "exit on command error (default false)")
}

func main() {
	flag.Parse()

	if strings.ToLower(flag.Arg(0)) == "version" {
		printVersion()
		os.Exit(0)
	}

	if err := validateWatcher(CommandWatcher); err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	os.Exit(CommandWatcher.Start())
}

func printVersion() {
	fmt.Printf("taildiff version %d.%d.%d\n", versionMajor, versionMinor, versionPatch)
}

func validateWatcher(w *watcher) error {
	if w.command == "" {
		return fmt.Errorf("A command is required. Supply command with -c <command>")
	}
	return nil
}
