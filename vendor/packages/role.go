package packages

import (
    "fmt"
    "os"
    "path/filepath"
    // "sort"
)

type Role struct {
    Name string
    Path string
    Files []File
    // Files map[string]File
    // Platform string
}

type VisitFunc func(string, os.FileInfo, *Role) error

func (r *Role) Explore(path string, fn VisitFunc) error {
    d, err := os.Open(path)
    if err != nil {
        return err
    }
    defer d.Close()
    di, err := d.Readdir(-1) // names
    if err != nil {
        return err
    }
    // sort.Strings(di)
    FILES:
    for _, fi := range di {
        switch filepath.Ext(fi.Name()) {
        case ".tpl", ".pkg":
            continue FILES
        }
        if fi.IsDir() {
            r.Explore(filepath.Join(path, fi.Name()), fn)
        } else {
            err := fn(path, fi, r)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func (r *Role) Sync() error {
    for _, f := range r.Files {
        if f.IsLinked() {
            fmt.Printf("%s: %s already linked!\n", r.Name+"\\"+f.Name)
        } else if f.IsLink() {
            fmt.Printf("%s: %s already linked to %s\n", r.Name+"\\"+f.Name, f.Dest, f.Link)
        } else if f.Exists() {
            fmt.Printf("%s: %s already exists\n", r.Name+"\\"+f.Name, f.Dest)
        } else {
            fmt.Printf("%s: %s does not exists!\n", r.Name+"\\"+f.Name, f.Dest)
        }
    }
    return nil
}
