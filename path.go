package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var ext = filepath.Ext

func exists(path string) bool {
    if _, err := os.Stat(path); err != nil {
        // if os.IsNotExist(err) {
        //     return false
        // }
        msg := strings.Replace(err.Error(), "stat ", os.Args[0]+": ", 1)
        fmt.Fprintf(os.Stderr, "%s\n", msg)
        return false
    }
    return true
}

func join(paths ...string) string {
    return filepath.Join(paths...)
}
