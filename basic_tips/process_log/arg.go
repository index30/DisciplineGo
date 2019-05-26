package main

import (
	"flag"
)

func parser() (string, string, []string) {
	logDir := flag.String(
		"logdir", "", "Output log directory (default=stderr)")
	flag.Parse()
	return *logDir, flag.Arg(0), flag.Args()[1:]
}