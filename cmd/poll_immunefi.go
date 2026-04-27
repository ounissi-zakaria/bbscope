package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ounissi-zakaria/bbscope/pkg/platforms"
	implatform "github.com/ounissi-zakaria/bbscope/pkg/platforms/immunefi"
	"github.com/ounissi-zakaria/bbscope/pkg/whttp"
)

// poll immunefi: shorthand for Immunefi
var pollImmunefiCmd = &cobra.Command{
	Use:   "immunefi",
	Short: "Poll Immunefi programs",
	RunE: func(cmd *cobra.Command, _ []string) error {
		proxy, _ := rootCmd.Flags().GetString("proxy")
		if proxy != "" {
			whttp.SetupProxy(proxy)
		}
		poller := &implatform.Poller{}
		return runPollWithPollers(cmd, []platforms.PlatformPoller{poller})
	},
}

func init() {
	pollCmd.AddCommand(pollImmunefiCmd)
}
