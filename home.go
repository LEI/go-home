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

var (
    debug bool
    verbose = 0
    src = ""
    dst = os.Getenv("HOME")
    act string
    ignoreDirs = []string{".git", "lib"}
    // ignore = []string{"*.tpl", ".pkg"}
    roles []string
)

func Usage(e int, msg ...interface{}) {
    if len(msg) > 0 {
        fmt.Fprintf(os.Stderr, "%s\n", msg...)
        // fmt.Fprintf(os.Stderr, "%s\n", optarg.UsageInfo)
    }
    optarg.Usage()
    os.Exit(e)
}

func main() {
    remain = getOpts()
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

func getOpts() []string {
    optarg.UsageInfo = fmt.Sprintf("Usage:\n\n  %s [options] [roles...]", os.Args[0]) // <action> hdvstIR
    optarg.HeaderFmt = "\n%s:"

    optarg.Header("General options")
    optarg.Add("h", "help", "Displays this help", false)
    optarg.Add("d", "debug", "Check mode", false)
    optarg.Add("v", "verbose", "Print more (default to: 0)", false)

    optarg.Header("Paths")
    optarg.Add("s", "source", "Source directory", src)
    optarg.Add("t", "target", "Target directory", dst)
    // optarg.Add("i", "ignore", "Exclude path", ignore)

    optarg.Header("Actions")
    optarg.Add("I", "install", "", true)
    optarg.Add("R", "remove", "", false)

    for opt := range optarg.Parse() {
        switch opt.ShortName {
            case "h":
                Usage(0)
            case "d":
                debug = opt.Bool()
            case "v":
                if opt.Bool() {
                    verbose += 1
                }

            case "s":
                src = opt.String()
                // Prompt?
            case "t":
                dst = opt.String()

            // case "I", "R":
            //     act = opt.String()
        }
    }

    if act == "" {
        Usage(1, "missing action: install or remove")
    }
    if !exists(src) {
        Usage(1, src + ": source directory does not exist")
    }
    if !exists(dst) {
        Usage(1, dst + ": destination directory does not exist")
    }

    return optarg.Remainder
}

func check(path string, info os.FileInfo, err error) error {
    // fmt.Printf("%s?\n", path)
    if err != nil {
        return err
    }
    if !info.IsDir() {
        // if verbose > 0 { fmt.Printf("Not a directory: %s\n", path) }
        return dir.Skip
    }
    name := info.Name()
    for _, i := range ignoreDirs {
        // if regexp.MustCompile(r).MatchString(p.Name())
        if name == i {
            // fmt.Println(name, "ignored")
            return dir.Skip
        }
    }
    if strings.HasPrefix(name, "os_") {
        if name == "os_" + OS {
            // fmt.Println("filepath.Walk", path, checkDir)
            err := walk(path)
            if err != nil {
                return err
            }
        }
        return dir.Skip
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
    d, err := dir.Read(path)
    if err != nil {
        return err
    }
    if verbose > 0  {
        fmt.Printf("DIR %s\n", join(path))
    }
    FILES:
    for _, fi := range d {
        switch filepath.Ext(fi.Name()) {
            case ".tpl", ".pkg":
                continue FILES
        }
        // for _, i := range ignore { }
        s := join(path, fi.Name())
        t := join(dst, fi.Name())
        if verbose > 0  {
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

func walk(path string) error {
    return dir.Walk(path, check, visit)
    // if err != nil {
    //     return err
    // }
    // return nil
}

func exists(path string) bool {
    if _, err := os.Stat(path); err != nil {
        if os.IsNotExist(err) {
            return false
        }
        fmt.Errorf("%s", err)
        return false
    }
    return true
}

func join(paths ...string) string {
    return filepath.Join(paths...)
}
