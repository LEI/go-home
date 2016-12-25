package main

import (
    "fmt"
    "log"
    "os"
    // "path/filepath"
    // "regexp"
    "runtime"
    "strings"
)

const OS = runtime.GOOS

var (
    ignoreDirs = []string{".git", "lib"}
    onlyDirs   []string
    // ignore = []string{"*.tpl", ".pkg"}
    // roles []string
)

func main() {
    onlyDirs = OptArg()
    // err := filepath.Walk(path, walkFn)
    err := walk(src)
    if err != nil {
        log.Fatal(err)
        // os.Exit(1)
    }
}

func found(dir string, fi os.FileInfo, lvl int) {
    s := join(dir, fi.Name())
    // t := join(path.Split(dir), ..., fi.Name())
    // t := strings.Replace(s, base, dst, 1)
    // t := strings.Replace(s, "/*$", "", -1)
    t := join(dst, fi.Name())
    if verbose > 0 {
        // fmt.Printf("%s <- %s\n", t, fi.Name())
        fmt.Printf("ln -s %s %s\n", s, t)
    }
    // err := os.Symlink(s, t)
    // if err != nil {
    //     return err
    // }
    // realpath, err := filepath.EvalSymlinks(t)
}

func walk(dir string) error {
    return WalkDir(dir, check, visit)
    // if err != nil {
    //     return err
    // }
    // return nil
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

func visit(dir string, info os.FileInfo, e error) error {
    if e != nil {
        return e
    }
    // if exists(join(dir, ".pkg")) {
    //     fmt.Println(dir, "READ PKG")
    // }
    // d := join(dir, info.Name())
    err := explore(dir, found, 0)
    if err != nil {
        return err
    }
    return nil
}

func explore(dir string, fn func(string, os.FileInfo, int), lvl int) error {
    if verbose > 0 {
        fmt.Printf("DIR %d  %s\n", lvl, dir)
    }
    d, err := ReadDir(dir)
    if err != nil {
        return err
    }
    FILES:
    for _, fi := range d {
        switch ext(fi.Name()) {
        case ".tpl", ".pkg":
            continue FILES
        }
        if fi.IsDir() {
            lvl++
            explore(join(dir, fi.Name()), fn, lvl)
        } else {
            fn(dir, fi, lvl)
        }
    }
    return nil
}
