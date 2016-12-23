package dir

import (
    "fmt"
    "os"
    "path/filepath"
    "runtime"
    "strings"
    "path/filepath"
    "dir"
)

var (
    Prefix = "os_"
    OS = runtime.GOOS
    Ignore []string
)

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
    for _, i := range Ignore {
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
