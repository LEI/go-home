package main

import (
    "fmt"
    "os"
    "github.com/jteeuwen/go-pkg-optarg"
    // "reflect"
)

// type OptArg interface {
//     Key() string
//     // Parse(interface{})
//     Set()
//     Type() string
// }

type Option struct {
    Name        string
    ShortName   string
    Description string
    defaultval  interface{}
    value       string
    // SetOpt(shortname string, name string, description string, defaultvalue interface{})
    // *optarg.Option
    parse       func(*optarg.Option)
}
func (o *Option) Key() string {
    return o.ShortName
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

func usage(e int, msg ...interface{}) {
    if len(msg) > 0 {
        fmt.Fprintf(os.Stderr, "%s\n", msg...)
        // fmt.Fprintf(os.Stderr, "%s\n", optarg.UsageInfo)
    }
    optarg.Usage()
    os.Exit(e)
}

func setOpts(opts []interface{}) map[string]interface{} {
    var oMap = make(map[string]interface{})
    for _, o := range opts {
        // fmt.Println(reflect.TypeOf(o))
        switch o.(type) {
            case *Option:
                opt := o.(*Option)
                oMap[opt.ShortName] = opt
                opt.Set()
            case *Header:
                opt := o.(*Header)
                opt.Set()
            // default:
            //     fmt.Println(o)
        }
    }
    return oMap
}

func getOpts(opts []interface{}) []string {
    optarg.UsageInfo = fmt.Sprintf("Usage:\n\n  %s [options] [roles...]", os.Args[0]) // <action> hdvstIR
    optarg.HeaderFmt = "\n%s:"

    oMap := setOpts(opts)
    for opt := range optarg.Parse() {
        oMap[opt.ShortName].(*Option).Parse(opt)
        // switch opt.ShortName {
        // case "h":
        //     usage(0)
        // case "d":
        //     debug = opt.Bool()
        // case "v":
        //     if opt.Bool() {
        //         verbose += 1
        //     }

        // case "s":
        //     src = opt.String()
        //     // Prompt?
        // case "t":
        //     dst = opt.String()

        // case "I", "R":
        //     act = opt.String()
        // }
    }

    return optarg.Remainder
}
