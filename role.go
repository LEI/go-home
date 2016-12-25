package main

import (
    "fmt"
)

type Role struct {
    Name string
    Files map[string]File
    // Platform string
}

func (r *Role) Play() error {
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
