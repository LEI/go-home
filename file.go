package main

import (
    "log"
    "os"
)

type File struct {
    Name string
    Dest string
    Link string
    Mode int64
    Stat os.FileInfo
    Source string
    didStat bool
    didRead bool
    exists bool
}

func (f *File) Exists() bool {
    if f.didStat != true {
        stat, err := os.Stat(f.Dest)
        f.exists = stat != nil
        if err != nil {
            if !os.IsNotExist(err) {
                log.Fatal(err)
            } else if os.IsExist(err) {
               f.exists = true
            }
        }
        f.didStat = true
    }
    return f.exists
}

func (f *File) IsLink() bool {
    if f.didRead != true {
        link, err := os.Readlink(f.Dest)
        f.Link = link
        if err != nil && f.exists == true {
            log.Fatal(err)
        }
        f.didRead = true
    }
    return len(f.Link) > 0
}

func (f *File) IsLinked() bool {
    return len(f.Link) > 0 && f.Link == f.Dest
}
