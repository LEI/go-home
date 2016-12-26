package main

type Config struct {
    OsPkg  map[string]OsPkg `toml:"os"`
    RolePkg map[string]RolePkg `toml:"pkg"`
}

type Pkg struct {
    List []string
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
    Tpl   map[string]Tpl
    Vars  []string
    // OsPkg
}

type Tpl struct {
    Prompt string
}
