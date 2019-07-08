package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "hugo plugin"
	app.Usage = "hugo plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "buildDrafts",
			Usage:  " include content marked as draft",
			EnvVar: "PLUGIN_BUILDDRAFTS",
		},
		cli.BoolFlag{
			Name:   "buildExpired",
			Usage:  "include expired content",
			EnvVar: "PLUGIN_BUILDEXPIRED",
		},
		cli.BoolFlag{
			Name:   "buildFuture",
			Usage:  "include content with publishdate in the future",
			EnvVar: "PLUGIN_BUILDFUTURE",
		},
		cli.StringFlag{
			Name:   "cacheDir",
			Usage:  "change cache directory (useful when using caching plugins)",
			EnvVar: "PLUGIN_CACHEDIR",
			Value:  "",
		},
		cli.StringFlag{
			Name:   "config",
			Usage:  "config file (default is path/config.yaml|json|toml)",
			EnvVar: "PLUGIN_CONFIG",
			Value:  "",
		},
		cli.StringFlag{
			Name:   "content",
			Usage:  "filesystem path to content directory",
			EnvVar: "PLUGIN_CONTENT",
			Value:  "",
		},
		cli.StringFlag{
			Name:   "layout",
			Usage:  "filesystem path to layout directory",
			EnvVar: "PLUGIN_LAYOUT",
			Value:  "",
		},
		cli.StringFlag{
			Name:   "output",
			Usage:  "filesystem path to write files to",
			EnvVar: "PLUGIN_OUTPUT",
			Value:  "",
		},
		cli.StringFlag{
			Name:   "source",
			Usage:  "filesystem path to read files relative from",
			EnvVar: "PLUGIN_SOURCE",
			Value:  "",
		},
		cli.StringFlag{
			Name:   "theme",
			Usage:  "theme to use (located in /themes/THEMENAME/)",
			EnvVar: "PLUGIN_THEME",
			Value:  "",
		},
		cli.StringFlag{
			Name:   "url",
			Usage:  "hostname (and path) to the root",
			EnvVar: "PLUGIN_URL",
			Value:  "",
		},
		cli.BoolFlag{
			Name:   "validate",
			Usage:  "validate config file before generation",
			EnvVar: "PLUGIN_VALIDATE",
		},
		cli.StringFlag{
			Name:   "hugoVersion",
			Usage:  "the hugo version to be used",
			EnvVar: "PLUGIN_HUGO_VERSION",
			Value:  "",
		},
		cli.BoolFlag{
			Name:   "hugoExtended",
			Usage:  "If the hugo extended package should be used",
			EnvVar: "PLUGIN_EXTENDED",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Config: Config{
			HugoVersion:  c.String("hugoVersion"),
			HugoExtended: c.Bool("hugoExtended"),
			BuildDrafts:  c.Bool("buildDrafts"),
			BuildExpired: c.Bool("buildExpired"),
			BuildFuture:  c.Bool("buildFuture"),
			Validate:     c.Bool("validate"),
			Config:       c.String("config"),
			Content:      c.String("content"),
			Layout:       c.String("layout"),
			Output:       c.String("output"),
			Source:       c.String("source"),
			Theme:        c.String("theme"),
			Url:          c.String("url"),
		},
		BuildInVersion: os.Getenv("PLUGIN_HUGO_SHIPPED_VERSION"),
	}
	return plugin.Exec()
}
