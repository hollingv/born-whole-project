package main

import (
	"os"
	"slices"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type statusItem struct {
	name        string
	description string
	status      string
}

func checkEnvVar(name string) string {
	if os.Getenv(name) != "" {
		return "OK"
	}
	return "X"
}

const (
	flagNameSilent      = "silent"
	flagNameSetExitCode = "set-exit-code"
	statusOK            = "OK"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display status of required environment variables",
	Long:  `Display the status of required environment variables for gacli`,
	Run: func(cmd *cobra.Command, args []string) {
		silent, _ := cmd.Flags().GetBool(flagNameSilent)
		setExitCode, _ := cmd.Flags().GetBool(flagNameSetExitCode)

		items := []statusItem{
			{projectPrefix + "_CF_API_TOKEN", "Cloudflare API token (for AI)", checkEnvVar(projectPrefix + "_CF_API_TOKEN")},
			{projectPrefix + "_CF_ACCOUNT_ID", "Cloudflare account ID (for AI)", checkEnvVar(projectPrefix + "_CF_ACCOUNT_ID")},
			{projectPrefix + "_YT_API_KEY", "YouTube Data API key (for Shorts)", checkEnvVar(projectPrefix + "_YT_API_KEY")},
		}

		if !silent {
			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"Variable", "Description", "Status"})
			for _, item := range items {
				table.Append([]string{item.name, item.description, item.status})
			}
			table.Render()
		}

		if setExitCode {
			allOK := !slices.ContainsFunc(items, func(i statusItem) bool {
				return i.status != statusOK
			})
			if !allOK {
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().Bool(flagNameSilent, false, "Suppress all output")
	statusCmd.Flags().Bool(flagNameSetExitCode, false, "Exit with code 1 if any items are not OK")
}
