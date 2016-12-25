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
    // ignore = []string{"*.tpl", ".pkg"}
    // roles []string
)

func walk(path string) error {
    return WalkDir(path, check, visit)
    // if err != nil {
    //     return err
    // }
    // return nil
}

func check(path string, info os.FileInfo, err error) error {
    // fmt.Printf("%s?\n", path)
    if err != nil {
        return err
    }
    if !info.IsDir() {
        // if verbose > 0 { fmt.Printf("Not a directory: %s\n", path) }
        return Skip
    }
    name := info.Name()
    for _, i := range ignoreDirs {
        // if regexp.MustCompile(r).MatchString(p.Name())
        if name == i {
            // fmt.Println(name, "ignored")
            return Skip
        }
    }
    if strings.HasPrefix(name, "os_") {
        if name == "os_"+OS {
            // fmt.Println("filepath.Walk", path, checkDir)
            err := walk(path)
            if err != nil {
                return err
            }
        }
        return Skip
    }
    return nil
}

func visit(path string, info os.FileInfo, e error) error {
    if e != nil {
        return e
    }
    // if exists(join(path, ".pkg")) {
    //     fmt.Println(path, "READ PKG")
    // }
    // d := join(path, info.Name())
    d, err := ReadDir(path)
    if err != nil {
        return err
    }
    if verbose > 0 {
        fmt.Printf("DIR %s\n", join(path))
    }
    FILES:
    for _, fi := range d {
        switch ext(fi.Name()) {
        case ".tpl", ".pkg":
            continue FILES
        }
        // for _, i := range ignore { }
        s := join(path, fi.Name())
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
    return nil
}

// type test func(p ...string)

func main() {
    remain := OptArg()
    // if debug {
    //     fmt.Printf("%s %s: %s -> %s\n", verbose, act, src, dst)
    // }
    // err := filepath.Walk(path, walkFn)
    if len(remain) > 0 {
        // Usage(1, "Extra arguments: " + strings.Join(remain, " "))
        for _, r := range remain {
            err := walk(join(src, r))
            if err != nil {
                log.Fatal(err)
            }
        }
    } else {
        err := walk(src)
        if err != nil {
            log.Fatal(err)
            // os.Exit(1)
        }
    }
}
