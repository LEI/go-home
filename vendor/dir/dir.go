package dir

import (
    "fmt"
    "os"
    "path/filepath"
)

var (
    IsDir = fmt.Errorf("this is a file")
    IsFile = fmt.Errorf("this is a directory")
    SkipDir = filepath.SkipDir
    SkipFile = fmt.Errorf("skip this file")
)


func Read(dirname string) ([]os.FileInfo, error) {
    f, err  := os.Open(dirname)
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

func Walk(path string, walkFn ...WalkFunc) error {
    p, err := Read(path)
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
                    case IsFile, IsDir, SkipDir, SkipFile:
                        continue DIRS
                }
                return err
            }
        }
    }
    return nil
}
