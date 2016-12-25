package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var (
    Skip     = fmt.Errorf("skip this path")
    NotFound = fmt.Errorf("no such file or directory")
)

type walkFunc func(path string, info os.FileInfo, err error) error

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

func readDir(dirname string) ([]os.FileInfo, error) {
    f, err := os.Open(dirname)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    paths, err := f.Readdir(-1) // names
    if err != nil {
        return nil, err
    }
    // sort.Strings(paths)
    return paths, nil
}

func walkDir(path string, walkFn ...walkFunc) error {
    p, err := readDir(path)
    if err != nil {
        return err
    }

    DIRS:
    for _, fi := range p {
        root := filepath.Join(path, fi.Name())
        // for _; fn := range walkFn { // unexpected range, expecting expression
        for i := 0; i < len(walkFn); i++ {
            err := walkFn[i](root, fi, nil)
            if err != nil {
                switch err {
                case Skip:
                    continue DIRS
                }
                return err
            }
        }
    }
    return nil
}
