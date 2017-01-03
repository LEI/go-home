package main

import (
    "fmt"
    "os"
    "github.com/jteeuwen/go-pkg-optarg"
    // "reflect"
)

type Option struct {
    Name        string
    ShortName   string
    Description string
    defaultval  interface{}
    value       string
    parse       func(*optarg.Option)
    // SetOpt(shortname string, name string, description string, defaultvalue interface{})
    // *optarg.Option
}

func (o *Option) Set() {
    optarg.Add(o.ShortName, o.Name, o.Description, o.defaultval)
}

func (o *Option) Parse(opt *optarg.Option) {
    o.parse(opt)
}

type Header struct {
    Text string
}
func (o *Header) Set() {
    optarg.Header(o.Text)
}

type Options map[string]interface{}

func parseOpts(oMap Options) []string {
    // oMap := setOpts(opts)
    for opt := range optarg.Parse() {
        // fmt.Println("parse", opt.ShortName)
        oMap[opt.ShortName].(*Option).Parse(opt)
    }
    return optarg.Remainder
}

func setOpts(opts []interface{}) Options {
    var oMap = make(Options)
    for _, o := range opts {
        switch o.(type) {
            case *Option:
                opt := o.(*Option)
                oMap[opt.ShortName] = opt
                opt.Set()
            case *Header:
                opt := o.(*Header)
                opt.Set()
            // default:
        }
    }
    return oMap
}

func usage(e int, msg ...interface{}) {
    if len(msg) > 0 {
        fmt.Fprintf(os.Stderr, "%s\n", msg...)
        // fmt.Fprintf(os.Stderr, "%s\n", optarg.UsageInfo)
    }
    optarg.Usage()
    os.Exit(e)
}
