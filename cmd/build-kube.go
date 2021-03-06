package cmd

import (
	"github.com/SUSE/fissile/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	flagBuildKubeOutputDir       string
	flagBuildKubeDefaultEnvFiles []string
	flagBuildKubeUseMemoryLimits bool
)

// buildKubeCmd represents the kube command
var buildKubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Creates Kubernetes configuration files.",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		flagBuildKubeOutputDir = buildKubeViper.GetString("output-dir")
		flagBuildKubeDefaultEnvFiles = splitNonEmpty(buildKubeViper.GetString("defaults-file"), ",")
		flagBuildKubeUseMemoryLimits = buildKubeViper.GetBool("use-memory-limits")

		err := fissile.LoadReleases(
			flagRelease,
			flagReleaseName,
			flagReleaseVersion,
			flagCacheDir,
		)
		if err != nil {
			return err
		}

		opinions, err := model.NewOpinions(
			flagLightOpinions,
			flagDarkOpinions,
		)
		if err != nil {
			return err
		}

		return fissile.GenerateKube(
			flagRoleManifest,
			flagBuildKubeOutputDir,
			flagRepository,
			flagDockerRegistry,
			flagDockerOrganization,
			fissile.Version,
			flagBuildKubeDefaultEnvFiles,
			flagBuildKubeUseMemoryLimits,
			false,
			"",
			opinions,
		)
	},
}
var buildKubeViper = viper.New()

func init() {
	initViper(buildKubeViper)

	buildCmd.AddCommand(buildKubeCmd)

	buildKubeCmd.PersistentFlags().StringP(
		"output-dir",
		"",
		".",
		"Kubernetes configuration files will be written to this directory",
	)

	buildKubeCmd.PersistentFlags().StringP(
		"defaults-file",
		"D",
		"",
		"Env files that contain defaults for the parameters generated by kube",
	)

	buildKubeCmd.PersistentFlags().BoolP(
		"use-memory-limits",
		"",
		true,
		"Include memory limits when generating kube configurations",
	)

	buildKubeViper.BindPFlags(buildKubeCmd.PersistentFlags())
}
