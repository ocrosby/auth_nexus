package cmd

import (
	"github.com/ocrosby/auth-nexus/internal/utils"
	"github.com/spf13/cobra"
)

// tidyAllCmd represents the tidy-all command
var tidyAllCmd = &cobra.Command{
	Use:   "tidy-all",
	Short: "Run 'go mod tidy' on all modules in the monorepo",
	Long:  `This command runs 'go mod tidy' for each Go module in the monorepo to ensure that their dependencies are in sync.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ListAndTidyServicesModules()
	},
}

func init() {
	rootCmd.AddCommand(tidyAllCmd)
}
