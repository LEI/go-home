package main

import (
    "log"
    "os"
    "github.com/jteeuwen/go-pkg-optarg"
    // "regexp"
    "runtime"
)

const OS = runtime.GOOS

var (
    debug   bool
    verbose = 0
    src     = ""
    dst     = os.Getenv("HOME")
    act     string
)

var opts = []interface{} {
        &Header{Text: "General Options"},
        &Option{
            ShortName: "h", Name: "help",
            Description: "Displays this help", defaultval: false,
            parse: func(opt *optarg.Option) {
                usage(0)
        }},
        &Option{
            ShortName: "d", Name: "debug",
            Description: "Check mode", defaultval: false,
            parse: func(opt *optarg.Option) {
                debug = opt.Bool()
        }},
        &Option{
            ShortName: "v", Name: "verbose",
            Description: "Print more (default to: 0)", defaultval: false,
            parse: func(opt *optarg.Option) {
                if opt.Bool() {
                    verbose += 1
                    // verbose += opt.Int()
                }
        }},

        &Header{Text: "Paths"},
        &Option{
            ShortName: "s", Name: "source",
            Description: "Source directory", defaultval: src,
            parse: func(opt *optarg.Option) {
                src = opt.String()
                if src == "" {
                    log.Fatal("missing source directory")
                }
        }},
        &Option{
            ShortName: "t", Name: "target",
            Description: "Target directory", defaultval: dst,
            parse: func(opt *optarg.Option) {
                dst = opt.String()
        }},

        &Header{Text: "Actions (default to: install)"},
        &Option{
            ShortName: "I", Name: "Install",
            Description: "", defaultval: true,
            parse: func(opt *optarg.Option) {
                if opt.Bool() { act = "install" }
        }},
        &Option{
            ShortName: "R", Name: "remove",
            Description: "", defaultval: false,
            parse: func(opt *optarg.Option) {
                if opt.Bool() { act = "remove" }
        }},
    }
