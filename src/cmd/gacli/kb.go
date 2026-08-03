package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var kbBuildCmd = &cobra.Command{
	Use:   "kb-build",
	Short: "Build the knowledge base",
	Long:  `Build the knowledge base`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Knowledge build...")
	},
}

func init() {
	rootCmd.AddCommand(kbBuildCmd)
}
