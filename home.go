package main

import (
    "path/filepath"
    "log"
    "os"
    "fmt"
    // "regexp"
    "runtime"
    "strings"
    "github.com/jteeuwen/go-pkg-optarg"
)

const OS = runtime.GOOS
var ignoreDirs = []string{".git", "lib"}

func initArgs() {
    optarg.Header("General")
    optarg.Add("h", "help", "Displays this help", false)
    optarg.Add("d", "debug", "Check mode", false)
    optarg.Add("v", "verbose", "Print more (default to: 0)", false)

    optarg.Header("Paths")
    optarg.Add("s", "source", "Source directory", "~/.dotfiles")
    optarg.Add("t", "target", "Target directory", "$HOME")

    optarg.Header("Actions")
    optarg.Add("I", "install", "", true)
    optarg.Add("R", "remove", "", false)
}

func main() {
    initArgs()
    // os.Args = []string{os.Args[0], "-s", "~/.dotfiles", "-t", "$HOME", "-d"}

    var debug bool
    var verbose=0
    var src, dst, action string

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

            case "I":
                action = opt.String()
            case "R":
                action = opt.String()
        }
    }

    // if len(optarg.Remainder) != 0 {
    //     fmt.Errorf("Parse(): incorrect number of remaining arguments. Expected 2. got %d", len(Remainder))
    // }

    if debug {
        fmt.Printf("%s %s: %s -> %s\n", verbose, action, src, dst)
    }
    // fmt.Printf("%+v\n", flag.Args())

    // walkDir(src, visit)
    err := walkDir(src, visitDir)
    die(err)
}

func readDir(dirname string) ([]os.FileInfo, error) {
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

func walkDir(path string, walkFn WalkFunc) error {
    // root = path
    // err := filepath.Walk(path, walkFn)
    // return err

    // d, err := os.Open(path)
    // die(err)
    // defer d.Close()
    // p, err := d.Readdir(-1)
    p, err := readDir(path)
    die(err)

    // DIRS:
    for _, fi := range p {
        root := filepath.Join(path, fi.Name())
        err := walkFn(root, fi, nil)
        if err != nil {
            switch err {
                case filepath.SkipDir:
                    // err = nil
                    break
                default:
                    return err
            }
        }
    }
    return nil
}

func visitDir(path string, info os.FileInfo, err error) error {
    // fmt.Printf("%s?\n", path)
    die(err)
    if !info.IsDir() {
        return nil // fmt.Errorf("Not a directory: %s", path)
    }
    // if root == path {
    //     return nil
    // }

    name := info.Name()
    for _, i := range ignoreDirs {
        // if regexp.MustCompile(r).MatchString(p.Name())
        if name == i {
            // fmt.Println(name, "ignored")
            return filepath.SkipDir // fmt.Errorf("Ignored directory: %s", path)
        }
    }
    if strings.HasPrefix(name, "os_") {
        if name == "os_" + OS {
            // fmt.Println("filepath.Walk", path, checkDir)
            fmt.Printf("Nested: %s\n", path)
            walkDir(path, visitDir)
        }
        return filepath.SkipDir // fmt.Errorf("Matched prefix: %s", name)
    }
    fmt.Printf("Visited: %s\n", path)
    // if strings.HasPrefix(name, prefix) {
    //     // FIXME name -> fullPath
    //     walkDir(name, prefix) // Recursive prefix, handle empty!
    // } else {
    //     fmt.Println(name)
    // }
    return nil
}

func die(err error) {
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}
