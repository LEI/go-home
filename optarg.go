package main

import (
    "fmt"
    "os"
    "github.com/jteeuwen/go-pkg-optarg"
)

var (
    debug   bool
    verbose = 0
    src     = ""
    dst     = os.Getenv("HOME")
    act     string
)

func Usage(e int, msg ...interface{}) {
    if len(msg) > 0 {
        fmt.Fprintf(os.Stderr, "%s\n", msg...)
        // fmt.Fprintf(os.Stderr, "%s\n", optarg.UsageInfo)
    }
    optarg.Usage()
    os.Exit(e)
}

func OptArg() []string {
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

        case "I", "R":
            act = opt.String()
        }
    }

    if act == "" {
        Usage(1, "missing action: install or remove")
    }
    if !exists(src) {
        Usage(1)
    }
    if !exists(dst) {
        Usage(1)
    }

    return optarg.Remainder
}
