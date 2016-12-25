package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    // "regexp"
    "strings"
)

var (
    // visited    = make(map[string]map[string]File)
    visited    = make(map[string]Role)
    ignoreDirs = []string{".git", "lib"}
    onlyDirs   []string
    // ignore = []string{"*.tpl", ".pkg"}
    // roles []string
)

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
        logFatal(err)
        // os.Exit(1)
    }
    if len(onlyDirs) > 0 && len(visited) == 0 {
        warn("%s/{%v}: no such role", src, strings.Join(onlyDirs, ","))
    } else if len(visited) == 0 {
        warn("%s: no such role", src)
    }
    for name, role := range visited {
        if verbose > 1 { fmt.Println("ROLE", strings.ToUpper(name)) }
        for _, f := range role.Files {
            if f.IsLinked() {
                fmt.Printf("%s: %s already linked!\n", name+"\\"+f.Name)
            } else if f.IsLink() {
                fmt.Printf("%s: %s already linked to %s\n", name+"\\"+f.Name, f.Dest, f.Link)
            } else if f.Exists() {
                fmt.Printf("%s: %s already exists\n", name+"\\"+f.Name, f.Dest)
            } else {
                fmt.Printf("%s: %s does not exists!\n", name+"\\"+f.Name, f.Dest)
            }
        }
    }
}

func logErr(err error, replace ...string) error {
    msg := err.Error()
    // if len(replace) == 0 {
    //     replace = []string{"stat "}
    // }
    for _, r := range replace {
        msg = strings.Replace(msg, r, os.Args[0]+": ", 1)
    }
    return fmt.Errorf("%s\n", msg)
}

func logFatal(err error, replace ...string) {
    log.Fatal(logErr(err, replace...))
}

func warn(f string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, f+"\n", args...)
}

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
    role := filepath.Base(dir)
    // if verbose > 0 {
    //     fmt.Printf("ROLE: %v\n", role)
    // }
    visited[role] = Role{Name: role, Files: make(map[string]File)}
    err = explore(dir, found, role)
    if err != nil {
        return err
    }
    return nil
}

type VisitFunc func(string, os.FileInfo, string) error

func explore(path string, fn VisitFunc, role string) error {
    d, err := readDir(path)
    if err != nil {
        return err
    }
    // if len(d) == 0 {
    //     fi, err := os.Stat(path)
    //     if err != nil {
    //         return err
    //     }
    //     err = fn(path, fi, role)
    //     if err != nil {
    //         return err
    //     }
    // }
    FILES:
    for _, fi := range d {
        switch filepath.Ext(fi.Name()) {
        case ".tpl", ".pkg":
            continue FILES
        }
        if fi.IsDir() {
            explore(filepath.Join(path, fi.Name()), fn, role)
        } else {
            err = fn(path, fi, role)
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
    visited[role].Files[s] = File{Name: name, Dest: t}
        // Link: link,
        // Mode: int64(tStat.Mode()),
        // Stat: tStat,
    // tStat, err := os.Stat(t)
    // if err != nil && !os.IsNotExist(err) {
    //     return logErr(err, "stat ")
    // }
    // if err == nil {
    //     // if link == t {
    //     //     owned = true
    //     // }
    // }
    // if tStat != nil {
    //     fmt.Printf(">> %v\n", link)
    // }

    // link := ""
    // if tStat != nil {
    //     link, err := filepath.EvalSymlinks(t)
    //     if err != nil {
    //         return err
    //     }
    //     if t != link {
    //         warn("%s: %s already linked to %s", os.Args[0], t, link)
    //     } else if t == link {
    //         warn("%s: %s already owned", os.Args[0], t)
    //     } else {
    //         warn("%s: %s already exists", os.Args[0], t)
    //     }
    // }

    // targetFile := File{name: name, dest: t}
    // fileMap[s]


    // err := os.Symlink(s, t)
    // if err != nil {
    //     return err
    // }
    // realpath, err := filepath.EvalSymlinks(t)
    return nil
}
