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
    // roles []string
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
    remain := getOpts()
    dir.Ignore = ignoreDirs // []string{}
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
