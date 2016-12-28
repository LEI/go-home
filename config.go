package main

import (
    "fmt"
    "github.com/BurntSushi/toml"
    // "regexp"
)

type Config struct {
    OsPkg  map[string]OsPkg `toml:"os"`
    RolePkg map[string]RolePkg `toml:"pkg"`
}

type Pkg struct {
    Packages []interface{}
}

type OsPkg struct {
    *Pkg
    Install string
    Remove string
    Update string
}

type RolePkg struct {
    *Pkg
    OsPkg map[string]OsPkg
    Vars  map[string]Var
    // OsPkg
}

type Var struct {
    Prompt string
    Default string
}

func getConfig(path string) (Config, error) {
    var cfg Config
    if !exists(path) {
        return cfg, fmt.Errorf("%s no such file", path)
    }
    if _, err := toml.DecodeFile(path, &cfg); err != nil {
        return cfg, ErrorReplace(err)
    }
    // fmt.Printf("%v\n", cfg)
    for name, os := range cfg.OsPkg {
        if name == OS {
            fmt.Printf("%s packages = %v\n", name, os.Packages)
        }
    }
    for name, pkg := range cfg.RolePkg {
        fmt.Printf("role %s\n", name)
        fmt.Printf("packages = %v\n", pkg.Packages)
        for a, b := range pkg.Vars {
            fmt.Printf("%v = %v\n", a, b)
        }
    }
    return cfg, nil
}
