package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/spf13/cobra"
)

const (
	tmplPath   = "templates/index.html.tmpl"
	orgTmpl    = "templates/organization.html.tmpl"
	outputPath = "site/index.html"
)

const siteName = "Global Autonomy"

type templateData struct {
	SiteName  string
	OrgGroups []OrgGroup
}

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "Generate index.html from templates and organization data",
	Long:  `Renders templates/index.html.tmpl with organization data and writes site/index.html`,
	Run: func(cmd *cobra.Command, args []string) {
		tmpl, err := template.ParseFiles(tmplPath, orgTmpl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing templates: %v\n", err)
			os.Exit(1)
		}

		f, err := os.Create(outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating %s: %v\n", outputPath, err)
			os.Exit(1)
		}
		defer f.Close()

		if err := tmpl.Execute(f, templateData{SiteName: siteName, OrgGroups: orgGroups}); err != nil {
			fmt.Fprintf(os.Stderr, "Error rendering template: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Generated %s\n", outputPath)
	},
}

func init() {
	rootCmd.AddCommand(htmlCmd)
}
