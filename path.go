package main

import (
    "fmt"
    "os"
    // "path/filepath"
    "strings"
)

func basePath(path string, n int) string {
    sep := string(os.PathSeparator) // strconv.Itoa()
    p := strings.Split(path, sep)
    return strings.Join(p[:len(p)-n], sep)
}

func exists(path string) bool {
    if _, err := os.Stat(path); err != nil {
        msg := strings.Replace(err.Error(), "stat ", os.Args[0]+": ", 1)
        fmt.Fprintf(os.Stderr, "%s\n", msg)
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}
