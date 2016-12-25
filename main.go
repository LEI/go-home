package main

import (
    "fmt"
    "log"
    "os"
    "github.com/jteeuwen/go-pkg-optarg"
    "path/filepath"
    // "regexp"
    "runtime"
    "strings"
)

const OS = runtime.GOOS

var (
    debug   bool
    verbose = 0
    src     = ""
    dst     = os.Getenv("HOME")
    act     string
)

var (
    ignoreDirs = []string{".git", "lib"}
    onlyDirs   []string
    // ignore = []string{"*.tpl", ".pkg"}
    // roles []string
)

var opts = []interface{} {
    &Header{Text: "General Options"},
    &Option{ShortName: "h", Name: "help", Description: "Displays this help", defaultval: false, parse: func(opt *optarg.Option) {
        usage(0)
    }},
    &Option{ShortName: "d", Name: "debug", Description: "Check mode", defaultval: false, parse: func(opt *optarg.Option) {
        debug = opt.Bool()
    }},
    &Option{ShortName: "v", Name: "verbose", Description: "Print more (default to: 0)", defaultval: false, parse: func(opt *optarg.Option) {
        if opt.Bool() {
            verbose += 1
        // } else if int? {
        //     verbose += opt.Int()
        }
    }},

    &Header{Text: "Paths"},
    &Option{ShortName: "s", Name: "source", Description: "Source directory", defaultval: src, parse: func(opt *optarg.Option) {
        src = opt.String()
    }},
    &Option{ShortName: "t", Name: "target", Description: "Target directory", defaultval: dst, parse: func(opt *optarg.Option) {
        dst = opt.String()
    }},
    // optarg.Add("i", "ignore", "Exclude path", ignore)

    &Header{Text: "Actions (default to: install)"},
    &Option{ShortName: "I", Name: "Install", Description: "", defaultval: true, parse: func(opt *optarg.Option) {
        act = opt.String()
    }},
    &Option{ShortName: "R", Name: "remove", Description: "", defaultval: false, parse: func(opt *optarg.Option) {
        act = opt.String()
    }},
}

func main() {
    onlyDirs = getOpts(opts)
    if act == "" {
        usage(1, "missing action: install or remove")
    }
    if !exists(src) {
        usage(1)
    }
    if !exists(dst) {
        usage(1)
    }
    // err := filepath.Walk(path, walkFn)
    err := walk(src)
    if err != nil {
        log.Fatal(err)
        // os.Exit(1)
    }
}

type VisitFunc func(string, os.FileInfo, string) error

func walk(dir string) error {
    if !filepath.IsAbs(dir) {
        return fmt.Errorf("%s is not absolute", dir)
    }
    return walkDir(dir, check, visit)
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
    // if exists(filepath.Join(dir, ".pkg")) {
    //     fmt.Println(dir, "READ PKG")
    // }
    // d := filepath.Join(dir, info.Name())
    err = explore(dir, found, filepath.Base(dir))
    if err != nil {
        return err
    }
    return nil
}

func explore(dir string, fn VisitFunc, role string) error {
    if verbose > 0 {
        fmt.Printf("DIR %v\n", dir)
    }
    d, err := readDir(dir)
    if err != nil {
        return err
    }
    FILES:
    for _, fi := range d {
        switch filepath.Ext(fi.Name()) {
        case ".tpl", ".pkg":
            continue FILES
        }
        if fi.IsDir() { // TODO check empty?
            explore(filepath.Join(dir, fi.Name()), fn, role)
        } else {
            err = fn(dir, fi, role)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func found(path string, fi os.FileInfo, role string) error {
    name := fi.Name()
    s := filepath.Join(path, name)
    // t := strings.Replace(s, basePath(dir, n), dst, 1)
    base := filepath.Join(src, role)
    t := filepath.Join(strings.Replace(path, base, dst, 1), name)
    // fmt.Println(filepath.Join(src, base))
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
