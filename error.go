package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

func newErr(err error, replace ...string) error {
    msg := err.Error()
    // if len(replace) == 0 { replace = []string{"stat "} }
    for _, r := range replace {
        msg = strings.Replace(msg, r, os.Args[0]+": ", 1)
    }
    return fmt.Errorf("%s\n", msg)
}

func logErr(f string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, f+"\n", args...)
}

func logFatal(err error, replace ...string) {
    log.Fatal(newErr(err, replace...))
}
