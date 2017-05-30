package cmd

import (
	"github.com/fatih/color"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"

	"go.mikenewswanger.com/kube-ci/kube-ci/jobs"
	"go.mikenewswanger.com/utilities/filesystem"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Prepare application for run
		// Pre-cache jobs

		// Start listenning for requests

		// This is all temporary and should be moved into tests instead
		var fs = filesystem.Filesystem{}
		var contents, err = fs.LoadFileIfExists("~/documents/projects/kube-ci/job.yml")
		if err != nil {
			panic(err)
		}
		color.Yellow(contents)
		var job *jobs.Job
		err = yaml.Unmarshal([]byte(contents), &job)
		if err != nil {
			panic(err)
		}
		var y, _ = yaml.Marshal(job)
		color.Green(string(y))
		job.Trigger(map[string]string{})
	},
}

func init() {
	serveCmd.Flags().StringVarP(&commandLineFlags.configDirectory, "config", "c", "", "Configuration Directory")

	RootCmd.AddCommand(serveCmd)
}
