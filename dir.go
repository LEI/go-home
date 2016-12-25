package main

import (
    "fmt"
    "os"
)

var (
    Skip     = fmt.Errorf("skip this path")
    NotFound = fmt.Errorf("no such file or directory")
)

func ReadDir(dirname string) ([]os.FileInfo, error) {
    f, err := os.Open(dirname)
    if err != nil {
        return nil, err
    }
    // defer?
    paths, err := f.Readdir(-1) // names
    f.Close()
    if err != nil {
        return nil, err
    }
    // sort.Strings(paths)
    return paths, nil
}

type WalkFunc func(path string, info os.FileInfo, err error) error

func WalkDir(path string, walkFn ...WalkFunc) error {
    p, err := ReadDir(path)
    if err != nil {
        return err
    }

    DIRS:
    for _, fi := range p {
        root := join(path, fi.Name())
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