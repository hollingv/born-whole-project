package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/spf13/cobra"
)

const siteName = "Global Autonomy"

type templateData struct {
	SiteName  string
	OrgGroups []OrgGroup
}

type page struct {
	tmplFiles []string
	output    string
}

var pages = []page{
	{
		tmplFiles: []string{"templates/index.html.tmpl", "templates/nav.html.tmpl"},
		output:    "site/index.html",
	},
	{
		tmplFiles: []string{"templates/mission.html.tmpl", "templates/nav.html.tmpl"},
		output:    "site/mission.html",
	},
	{
		tmplFiles: []string{"templates/about.html.tmpl", "templates/nav.html.tmpl"},
		output:    "site/about.html",
	},
	{
		tmplFiles: []string{"templates/organizations.html.tmpl", "templates/nav.html.tmpl", "templates/organization.html.tmpl"},
		output:    "site/organizations.html",
	},
}

var siteCmd = &cobra.Command{
	Use:   "site",
	Short: "Generate site HTML from templates and organization data",
	Long:  `Renders all page templates and writes output to site/`,
	Run: func(cmd *cobra.Command, args []string) {
		data := templateData{SiteName: siteName, OrgGroups: orgGroups}

		for _, p := range pages {
			tmpl, err := template.ParseFiles(p.tmplFiles...)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing templates %v: %v\n", p.tmplFiles, err)
				os.Exit(1)
			}

			f, err := os.Create(p.output)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error creating %s: %v\n", p.output, err)
				os.Exit(1)
			}

			if err := tmpl.Execute(f, data); err != nil {
				f.Close()
				fmt.Fprintf(os.Stderr, "Error rendering %s: %v\n", p.output, err)
				os.Exit(1)
			}
			f.Close()
			fmt.Printf("Generated %s\n", p.output)
		}
	},
}

func init() {
	rootCmd.AddCommand(siteCmd)
}
