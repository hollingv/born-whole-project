package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"iwebsite/src/cmd/bictl/internal/data"

	"github.com/spf13/cobra"
)

type templateData struct {
	SiteName          string
	Version           string
	EnvPrefix         string
	OrgGroups         []data.OrgGroup
	FAQItems          []data.FAQItem
	ContributeExample string
	Resources         []data.Resource
}

type page struct {
	tmplFiles []string
	output    string
}

// discoverPages scans the templates directory and builds the list of pages to
// generate. Files containing {{define are treated as shared templates and
// included with every page. All other .tmpl files are treated as page templates.
func discoverPages(tmplDir string) ([]page, error) {
	entries, err := os.ReadDir(tmplDir)
	if err != nil {
		return nil, fmt.Errorf("reading templates directory: %w", err)
	}

	var shared []string
	var pageTmpls []string

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".tmpl") {
			continue
		}
		path := filepath.Join(tmplDir, e.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", path, err)
		}
		if strings.Contains(string(content), "{{define") {
			shared = append(shared, path)
		} else {
			pageTmpls = append(pageTmpls, path)
		}
	}

	var pages []page
	for _, pt := range pageTmpls {
		name := strings.TrimSuffix(filepath.Base(pt), ".tmpl")
		pages = append(pages, page{
			tmplFiles: append([]string{pt}, shared...),
			output:    filepath.Join("site", name),
		})
	}

	return pages, nil
}

var siteCmd = &cobra.Command{
	Use:   "site",
	Short: "Generate site HTML from templates and organization data",
	Long:  `Discovers all page templates and renders them to site/`,
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetString("version")
		tmplData := templateData{
			SiteName:          siteName,
			Version:           version,
			EnvPrefix:         projectPrefix,
			OrgGroups:         data.OrgGroups,
			FAQItems:          data.FAQItems,
			ContributeExample: data.ContributeExample,
			Resources:         data.LoadResources(),
		}

		pages, err := discoverPages("templates")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error discovering templates: %v\n", err)
			os.Exit(1)
		}

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

			if err := tmpl.Execute(f, tmplData); err != nil {
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
	siteCmd.Flags().String("version", "local-dirty", "Version string to embed in the page footer")
}
