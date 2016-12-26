package main

import (
    "fmt"
    "os"
    "path/filepath"
    "packages"
    // "regexp"
    "strings"
)

var (
    // visited    = make(map[string]map[string]packages.File)
    visited    = make(map[string]packages.Role)
    ignoreDirs = []string{".git", "lib"}
    onlyDirs   []string
    // ignore = []string{"*.tpl", ".pkg"}
    // roles []string
)

func main() {
    onlyDirs = getOpts(opts)
    switch false {
    case exists(src), exists(dst):
        usage(1)
    }
    switch act {
    // case "install"
    // case "remove"
    case "":
        usage(1, "missing action: install or remove")
    }
    // err := filepath.Walk(path, walkFn)
    err := walk(src)
    if err != nil {
        logFatal(err)
        // os.Exit(1)
    }
    if len(onlyDirs) > 0 && len(visited) == 0 {
        logErr("%s/{%v}: no such role", src, strings.Join(onlyDirs, ","))
    } else if len(visited) == 0 {
        logErr("%s: no such role", src)
    }
    for _, role := range visited {
        if verbose > 1 { fmt.Println("ROLE", strings.ToUpper(role.Name)) }
        switch act {
        case "install":
            err := role.Sync()
            if err != nil {
                logFatal(err)
            }
        }
    }
}

func walk(dir string) error {
    if !filepath.IsAbs(dir) {
        return fmt.Errorf("%s is not absolute", dir)
    }
    return walkDir(dir, check, visit)
}

func check(dir string, info os.FileInfo, err error) error {
    // fmt.Printf("%s?\n", dir)
    if err != nil {
        return err
    }
    if !info.IsDir() {
        // if verbose > 0 { fmt.Printf("Not a directory: %s\n", dir) }
        return Skip
    }
    name := info.Name()
    // if regexp.MustCompile(r).MatchString(p.Name())<Paste>
    for _, i := range ignoreDirs {
        if name == i { // fmt.Println(name, "ignored")
            return Skip
        }
    }
    if strings.HasPrefix(name, "os_") {
        if name == "os_"+OS {
            // fmt.Println("filepath.Walk", dir, checkDir)
            err := walk(dir)
            if err != nil {
                return err
            }
        }
        return Skip
    }
    for _, i := range onlyDirs {
        if name != i {
            return Skip
        }
    }
    return nil
}

func visit(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    role := packages.Role{
        Name: filepath.Base(path),
        Path: path,
        Files: []packages.File{},
    }
    // if verbose > 0 { fmt.Printf("ROLE: %v\n", role) }
    err = role.Explore(path, found)
    if err != nil {
        return err
    }
    visited = append(visited, role)
    return nil
}

func found(path string, fi os.FileInfo, role *packages.Role) error {
    role.Files = append(role.Files, packages.File{
        Name: fi.Name(),
        Dest: filepath.Join(strings.Replace(path, role.Path, dst, 1), fi.Name()),
        Source: filepath.Join(path, fi.Name()),
    })
    return nil
}
