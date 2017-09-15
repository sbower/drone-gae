package main

import (
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

var reRedact = regexp.MustCompile(`(oauth2_access_token)\s+\S+\s+`)

type Environ struct {
	dir    string
	env    []string
	stdout io.Writer
	stderr io.Writer
}

func NewEnviron(dir string, env []string, stdout, stderr io.Writer) *Environ {
	return &Environ{
		dir:    dir,
		env:    env,
		stdout: stdout,
		stderr: stderr,
	}
}

// Run executes the given program.
func (e *Environ) Run(name string, arg ...string) error {
	displayArg := reRedact.ReplaceAll([]byte(strings.Trim(fmt.Sprint(arg), "[]")), []byte("$1 [redacted] "))
	fmt.Printf("Running Command: %s %s\n", name, displayArg)
	cmd := exec.Command(name, arg...)
	cmd.Dir = e.dir
	cmd.Env = e.env
	cmd.Stdout = e.stdout
	cmd.Stderr = e.stderr
	return cmd.Run()
}
