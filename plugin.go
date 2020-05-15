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
		URL      string
		Drafts   bool
		Expired  bool
		Future   bool
		Validate bool
		Cache    string
		Config   string
		Content  string
		Layout   string
		Output   string
		Source   string
		Theme    string
		Version  string
		Extended bool
	}
)

var (
	hugoExecutable = "hugo"
)

// Exec executes the plugins functionality
func (p Plugin) Exec() error {
	var cmds = make([]*exec.Cmd, 0)

	if p.Config.Extended {
		hugoExecutable = "hugo-extended"
	}

	// Check if buildIn plugin version equals
	// plugin version declared in drone.yml
	if !versionsEqual(p.BuildInVersion, p.Config.Version) {
		version := p.Config.Version

		if version == "" {
			version = p.BuildInVersion
		}

		executable, err := download.Get(version, p.Config.Extended)

		if err != nil {
			return err
		}

		hugoExecutable = executable
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

	if config.Drafts {
		args = append(args, "-D")
	}

	if config.Expired {
		args = append(args, "-E")
	}

	if config.Future {
		args = append(args, "-F")
	}

	if config.Cache != "" {
		args = append(args, "--cacheDir", config.Cache)
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

	if config.URL != "" {
		args = append(args, "--baseURL", config.URL)
	}

	return exec.Command(hugoExecutable, args...)
}

// execAll executes a slice of commands as a batch job
func execAll(cmds []*exec.Cmd) error {
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

func versionsEqual(version string, toCompare string) bool {
	if toCompare == version || toCompare == "" {
		return true
	}

	return false
}

// trace writes each command to stdout with the command wrapped in an xml
// tag so that it can be extracted and displayed in the logs.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}
