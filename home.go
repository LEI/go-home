package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    // "regexp"
    "runtime"
    "strings"
    "github.com/jteeuwen/go-pkg-optarg"
    "dir"
)

// type Walk interface {
// }

const OS = runtime.GOOS
var ignoreDirs = []string{".git", "lib"}
// var sequence []dir.WalkFunc // []interface

var (
    debug bool
    verbose = 0
    src = ""
    dst = os.Getenv("HOME")
    action string
)

func parseArgs() {
    // args := {}
    optarg.Header("General")
    optarg.Add("h", "help", "Displays this help", false)
    optarg.Add("d", "debug", "Check mode", false)
    optarg.Add("v", "verbose", "Print more (default to: 0)", false)

    optarg.Header("Paths")
    optarg.Add("s", "source", "Source directory", src)
    optarg.Add("t", "target", "Target directory", dst)

    optarg.Header("Actions")
    optarg.Add("I", "install", "", true)
    optarg.Add("R", "remove", "", false)

    for opt := range optarg.Parse() {
        switch opt.ShortName {
            case "h":
                fmt.Printf("Help: %s", opt.Bool())
                optarg.Usage()
                os.Exit(0)
            case "d":
                debug = opt.Bool()
            case "v":
                if opt.Bool() {
                    verbose += 1
                }

            case "s":
                src = opt.String()
            case "t":
                dst = opt.String()

            case "I", "R":
                action = opt.String()
        }
    }
    // return args
}

func check(path string, info os.FileInfo, err error) error {
    // fmt.Printf("%s?\n", path)
    if err != nil {
        return err
    }
    if !info.IsDir() {
        // if verbose > 0 { fmt.Printf("Not a directory: %s\n", path) }
        return dir.IsFile
    }
    name := info.Name()
    for _, i := range ignoreDirs {
        // if regexp.MustCompile(r).MatchString(p.Name())
        if name == i {
            // fmt.Println(name, "ignored")
            return dir.SkipDir
        }
    }
    if strings.HasPrefix(name, "os_") {
        if name == "os_" + OS {
            // fmt.Println("filepath.Walk", path, checkDir)
            fmt.Printf("Nested: %s\n", path)
            err := walk(path)
            if err != nil {
                return err
            }
        }
        return dir.SkipDir
    }
    return nil
}

func visit(path string, info os.FileInfo, e error) error {
    if e != nil {
        return e
    }
    // d := join(path, info.Name())
    d, err := dir.Read(path)
    if err != nil {
        return err
    }
    fmt.Printf("DIR %s -> %s\n", join(path), dst)
    for _, fi := range d {
        // TODO ignore templates & cie
        s := join(path, fi.Name())
        t := join(dst, fi.Name())
        err := link(s, t)
        if err != nil {
            return err
        }
    }
    return nil
}

func link(source string, target string) error {
    fmt.Printf("Link %s to %s\n", source, target)
    return nil
}

func walk(path string) error {
    return dir.Walk(path, check, visit)
    // if err != nil {
    //     return err
    // }
    // return nil
}

func join(paths ...string) string {
    return filepath.Join(paths...)
}

func main() {
    parseArgs()

    // dir.OS = OS
    // dir.ignoreDirs = []string{".git", "lib"}
    // os.Args = []string{os.Args[0], "-s", "~/.dotfiles", "-t", "$HOME", "-d"}

    // if debug {
    //     fmt.Printf("%s %s: %s -> %s\n", verbose, action, src, dst)
    // }
    // fmt.Printf("%+v\n", flag.Args())

    // err := filepath.Walk(path, walkFn)
    err := walk(src)
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}
