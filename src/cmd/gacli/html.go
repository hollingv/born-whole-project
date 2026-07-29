package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const (
	orgsPath   = "content/organizations.yaml"
	tmplPath   = "templates/index.html.tmpl"
	orgTmpl    = "templates/organization.html.tmpl"
	outputPath = "site/index.html"
)

type Organization struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Website     string `yaml:"url"`
	Thumbnail   string `yaml:"thumbnail"`
}

type yamlData struct {
	Organizations []Organization `yaml:"organizations"`
}

type templateData struct {
	Organizations []Organization
}

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "Generate index.html from templates and organization data",
	Long:  `Reads site/index.html.tmpl and content/organizations.yaml, renders the page, and writes site/index.html`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load organizations from YAML
		raw, err := os.ReadFile(orgsPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", orgsPath, err)
			os.Exit(1)
		}
		var data yamlData
		if err := yaml.Unmarshal(raw, &data); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", orgsPath, err)
			os.Exit(1)
		}

		// Parse both templates together
		tmpl, err := template.ParseFiles(tmplPath, orgTmpl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing templates: %v\n", err)
			os.Exit(1)
		}

		// Write rendered output
		f, err := os.Create(outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating %s: %v\n", outputPath, err)
			os.Exit(1)
		}
		defer f.Close()

		if err := tmpl.Execute(f, templateData{Organizations: data.Organizations}); err != nil {
			fmt.Fprintf(os.Stderr, "Error rendering template: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Generated %s\n", outputPath)
	},
}

func init() {
	rootCmd.AddCommand(htmlCmd)
}
