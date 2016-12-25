package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
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

type VisitFunc func(string, os.FileInfo, int) error

func found(dir string, fi os.FileInfo, lvl int) error {
    // fmt.Printf("%v\n", filepath.SplitList(dir))

    // base := dir
    // for i := 0; i < lvl; i++ {
    //     base, _ := filepath.Split(base)
    //     // strings.Replace(base, "/*$", "", -1)
    //     if base == "" {
    //         return fmt.Errorf("invalid base path")
    //     }
    // }
    // fmt.Println(base, lvl)

    s := join(dir, fi.Name())
    // t := join(path.Split(dir), ..., fi.Name())
    // t := strings.Replace(s, base, dst, 1)
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
    return nil
}

func trimPath(p string, n int) (string, error) {
    return p, nil
}

func walk(dir string) error {
    if !filepath.IsAbs(dir) {
        return fmt.Errorf("%s is not absolute", dir)
    }
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

func visit(dir string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    // if exists(join(dir, ".pkg")) {
    //     fmt.Println(dir, "READ PKG")
    // }
    // d := join(dir, info.Name())
    err = explore(dir, found, 0)
    if err != nil {
        return err
    }
    return nil
}

func explore(dir string, fn VisitFunc, lvl int) error {
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
        if fi.IsDir() { // TODO check empty?
            lvl++
            explore(join(dir, fi.Name()), fn, lvl)
        } else {
            err = fn(dir, fi, lvl)
            if err != nil {
                return err
            }
        }
    }
    return nil
}
