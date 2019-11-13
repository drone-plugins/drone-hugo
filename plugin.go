package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/drone-plugins/drone-hugo/download"
)

type (
	Plugin struct {
		Config         Config
		BuildInVersion string
	}

	Config struct {
		BuildDrafts  bool
		BuildExpired bool
		BuildFuture  bool
		CacheDir     string
		Config       string
		Content      string
		Layout       string
		Output       string
		Source       string
		Theme        string
		Url          string
		HugoVersion  string
		HugoExtended bool
		Validate     bool
	}
)

var hugoExecutable = "hugo"

// Exec executes the plugins functionality
func (p Plugin) Exec() error {
	var cmds = make([]*exec.Cmd, 0)

	// Check if buildIn plugin version equals
	// plugin version declared in drone.yml
	if !versionsEqual(p.BuildInVersion, p.Config.HugoVersion, p.Config.HugoExtended) {
		hugoVersion := p.Config.HugoVersion
		if hugoVersion == "" {
			hugoVersion = p.BuildInVersion
		}

		hugoPath, err := download.Get(hugoVersion, p.Config.HugoExtended)
		if err != nil {
			return err
		}
		hugoExecutable = hugoPath
	}
	if p.Config.Validate {
		cmds = append(cmds, commandValidate(p.Config))
	}

	cmds = append(cmds, commandBuild(p.Config))
	return execAll(cmds)
}

func commandValidate(config Config) *exec.Cmd {
	args := []string{"check"}
	if config.Config != "" {
		args = append(args, "--config", config.Config)
	}
	return exec.Command(hugoExecutable, args...)
}

func commandBuild(config Config) *exec.Cmd {
	var args = make([]string, 0)

	// add bool args
	if config.BuildDrafts {
		args = append(args, "-D")
	}
	if config.BuildExpired {
		args = append(args, "-E")
	}
	if config.BuildFuture {
		args = append(args, "-F")
	}
	// add string args
	if config.CacheDir != "" {
		args = append(args, "--cacheDir", config.CacheDir)
	}
	if config.Config != "" {
		args = append(args, "--config", config.Config)
	}
	if config.Content != "" {
		args = append(args, "--contentDir", config.Content)
	}
	if config.Layout != "" {
		args = append(args, "--layoutDir", config.Layout)
	}
	if config.Output != "" {
		args = append(args, "--destination", config.Output)
	}
	if config.Source != "" {
		args = append(args, "--source", config.Source)
	}
	if config.Theme != "" {
		args = append(args, "--theme", config.Theme)
	}
	if config.Url != "" {
		args = append(args, "--baseURL", config.Url)
	}

	return exec.Command(hugoExecutable, args...)
}

// trace writes each command to stdout with the command wrapped in an xml
// tag so that it can be extracted and displayed in the logs.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}

func versionsEqual(version string, toCompare string, extended bool) bool {
	if extended {
		return false
	}

	if toCompare == version || toCompare == "" {
		return true
	} else {
		return false
	}
}

// execAll executes a slice of commands as a batch job
func execAll(cmds []*exec.Cmd) error {
	// Execute all commands
	for _, cmd := range cmds {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		trace(cmd)

		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
