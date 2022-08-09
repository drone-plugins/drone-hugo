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
			Name:   "drafts",
			Usage:  " include content marked as draft",
			EnvVar: "PLUGIN_BUILDDRAFTS,PLUGIN_DRAFTS",
		},
		cli.BoolFlag{
			Name:   "expired",
			Usage:  "include expired content",
			EnvVar: "PLUGIN_BUILDEXPIRED,PLUGIN_EXPIRED",
		},
		cli.BoolFlag{
			Name:   "future",
			Usage:  "include content with publishdate in the future",
			EnvVar: "PLUGIN_BUILDFUTURE,PLUGIN_FUTURE",
		},
		cli.StringFlag{
			Name:   "cache",
			Usage:  "change cache directory (useful when using caching plugins)",
			EnvVar: "PLUGIN_CACHEDIR,PLUGIN_CACHE",
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
		cli.StringFlag{
			Name:   "hugoversion",
			Usage:  "the hugo version to be used",
			EnvVar: "PLUGIN_HUGO_VERSION,PLUGIN_VERSION",
			Value:  "",
		},
		cli.BoolFlag{
			Name:   "extended",
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
			URL:      c.String("url"),
			Drafts:   c.Bool("drafts"),
			Expired:  c.Bool("expired"),
			Future:   c.Bool("future"),
			Config:   c.String("config"),
			Content:  c.String("content"),
			Layout:   c.String("layout"),
			Output:   c.String("output"),
			Source:   c.String("source"),
			Theme:    c.String("theme"),
			Version:  c.String("hugoversion"),
			Extended: c.Bool("extended"),
		},
		BuildInVersion: os.Getenv("PLUGIN_HUGO_SHIPPED_VERSION"),
	}

	return plugin.Exec()
}
