package app

import (
	"github.com/codegangsta/cli"
)

// CommandRouter will dispatch CLI commands to their relevant functions
func CommandRouter(c *cli.Context) {
	switch {
	case c.Command.FullName() == "release download":
	case c.Command.FullName() == "release jobs-report":
		ListJobs(c.String("release"))
	case c.Command.FullName() == "release packages-report":
		ListPackages(c.String("release"))
	case c.Command.FullName() == "compilation build-base":
		CreateBaseCompilationImage(
			c.String("base-image"),
			c.String("release"),
			c.String("repository"),
		)
	case c.Command.FullName() == "compilation show-base":
		ShowBaseImage(c.String("base-image"))
	case c.Command.FullName() == "compilation start":
		Compile(
			c.String("base-image"),
			c.String("release"),
			c.String("repository"),
			c.String("target"),
			c.Int("workers"),
		)
	case c.Command.FullName() == "configuration report":
		ListFullConfiguration(c.String("release"))
	case c.Command.FullName() == "templates report":
		PrintTemplateReport(c.String("release"))
	case c.Command.FullName() == "configuration generate":
		GenerateConfigurationBase(
			c.String("release"),
			c.String("light-opinions"),
			c.String("dark-opinions"),
			c.String("target"),
			c.String("prefix"),
			c.String("provider"),
		)
	}
}