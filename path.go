package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func basePath(path string, n int) string {
    sep := string(os.PathSeparator) // strconv.Itoa()
    p := strings.Split(path, sep)
    n = len(p) - n
    return strings.Join(p[:n], sep)
}

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
