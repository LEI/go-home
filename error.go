package main

import (
    "fmt"
    // "log"
    "os"
    "strings"
)

// type Error interface {
//     // Error() string
//     String() string
// }

// fmt.Fprintf(os.Stderr, f+"\n", args...)
// log.Fatal()

func ErrorReplace(err error, replace ...string) error {
    return NewError(err.Error(), replace...)
}

func NewError(err string, replace ...string) error {
    // if len(replace) == 0 { replace = []string{"stat "} }
    for _, r := range replace {
        err = strings.Replace(err, r, os.Args[0]+": ", 1)
    }
    return fmt.Errorf("%s\n", err)
}
