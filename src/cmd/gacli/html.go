package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "Generate HTML output",
	Long:  `Generate HTML output`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello from gacli")
	},
}

func init() {
	rootCmd.AddCommand(htmlCmd)
}
