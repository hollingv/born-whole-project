package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"born-whole-project/src/cmd/bwctl/internal/data"

	"github.com/spf13/cobra"
)

type templateData struct {
	RepoName          string
	SiteName          string
	SiteUrl           string
	ContactEmail      string
	Version           string
	ProjectPrefix     string
	OrgGroups         []data.OrgGroup
	FAQItems          []data.FAQItem
	SiteFAQItems      []data.FAQItem
	ContributeExample string
	Resources         []data.Resource
	Events            []data.Event
	NewsItems         []data.NewsItem
	DonationItems     []data.DonationItem
}

// EventGroup is a slice of events sharing the same Kind, used for template rendering.
type EventGroup struct {
	Kind   data.EventKind
	Label  string
	Events []data.Event
}

// HomeEventGroup holds an organisation and its capped list of events for the home page.
type HomeEventGroup struct {
	Org    *data.Organization
	Events []data.Event
}

// homeEventGroups prepares the home page event list: events are grouped by
// organisation (sorted by org name), with at most 2 events per org.
// The input is assumed to already be date-sorted, so the first 2 kept per org
// are the earliest upcoming.
func homeEventGroups(events []data.Event) []HomeEventGroup {
	var order []string
	seen := map[string]bool{}
	byOrg := map[string]*HomeEventGroup{}

	for _, e := range events {
		if !seen[e.OrgID] {
			seen[e.OrgID] = true
			order = append(order, e.OrgID)
			byOrg[e.OrgID] = &HomeEventGroup{Org: e.Org()}
		}
		grp := byOrg[e.OrgID]
		if len(grp.Events) < 2 {
			grp.Events = append(grp.Events, e)
		}
	}

	sort.Slice(order, func(i, j int) bool {
		ni, nj := order[i], order[j]
		if byOrg[ni].Org != nil {
			ni = byOrg[ni].Org.Name
		}
		if byOrg[nj].Org != nil {
			nj = byOrg[nj].Org.Name
		}
		return ni < nj
	})

	result := make([]HomeEventGroup, 0, len(order))
	for _, id := range order {
		result = append(result, *byOrg[id])
	}
	return result
}

// OrgSocialEntry pairs an organisation with its URL for a specific social platform.
type OrgSocialEntry struct {
	Org data.Organization
	URL string
}

// orgsBySocial returns all organisations across all groups that have the given
// social platform set, paired with their URL for that platform.
// platform must be "Facebook", "X", "Instagram", or "YouTube".
func orgsBySocial(groups []data.OrgGroup, platform string) []OrgSocialEntry {
	var entries []OrgSocialEntry
	for _, g := range groups {
		for _, org := range g.Organizations {
			var url string
			switch platform {
			case "Facebook":
				url = org.SocialSites.Facebook
			case "X":
				url = org.SocialSites.X
			case "Instagram":
				url = org.SocialSites.Instagram
			case "YouTube":
				url = org.SocialSites.YouTube
			}
			if url != "" {
				entries = append(entries, OrgSocialEntry{Org: org, URL: url})
			}
		}
	}
	return entries
}

// writeAttyClopperManifest writes the list of Attorney Clopper transcript files
// to dist/harvested/generated-attyclopper-kb-file-list.json so the
// /ask-attorney-clopper endpoint can search them. Paths are stored relative
// to dist/ (e.g. "kb/attyClopper/dQw4w9WgXcQ.txt").
func writeAttyClopperManifest() {
	files, _ := filepath.Glob(filepath.Join(attyClopperKBDir, "*.txt"))
	var paths []string
	for _, f := range files {
		paths = append(paths, "kb/attyClopper/"+filepath.Base(f))
	}
	raw, _ := json.MarshalIndent(paths, "", "  ")
	manifestPath := filepath.Join(targetDir, "harvested", "generated-attyclopper-kb-file-list.json")
	if err := os.WriteFile(manifestPath, raw, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing Attorney Clopper manifest: %v\n", err)
		return
	}
	fmt.Printf("Written %s (%d files)\n", manifestPath, len(paths))
}

// groupEventsByKind partitions events into groups by Kind, preserving the order
// in which each Kind is first encountered. It is registered as a template function.
func groupEventsByKind(events []data.Event) []EventGroup {
	var order []data.EventKind
	seen := map[data.EventKind]bool{}
	byKind := map[data.EventKind][]data.Event{}

	for _, e := range events {
		if !seen[e.Kind] {
			seen[e.Kind] = true
			order = append(order, e.Kind)
		}
		byKind[e.Kind] = append(byKind[e.Kind], e)
	}

	groups := make([]EventGroup, 0, len(order))
	for _, k := range order {
		groups = append(groups, EventGroup{
			Kind:   k,
			Label:  k.Label(),
			Events: byKind[k],
		})
	}
	return groups
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
			output:    filepath.Join(targetDir, name),
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
			SiteUrl:           siteUrl,
			ContactEmail:      contactEmail,
			RepoName:          repoName,
			Version:           version,
			ProjectPrefix:     projectPrefix,
			OrgGroups:         data.OrgGroups,
			FAQItems:          data.FAQItems,
			SiteFAQItems:      data.SiteFAQItems,
			ContributeExample: data.ContributeExample,
			Resources:         data.LoadResources(),
			Events:            data.AllEvents(),
			NewsItems:         data.NewsItems,
			DonationItems:     data.DonationItems,
		}

		// Ensure dist/harvested/ exists and generate the KB manifest.
		if err := os.MkdirAll(filepath.Join(targetDir, "harvested"), 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating harvested directory: %v\n", err)
			os.Exit(1)
		}
		writeManifest()
		writeAttyClopperManifest()

		pages, err := discoverPages("templates")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error discovering templates: %v\n", err)
			os.Exit(1)
		}

		funcMap := template.FuncMap{
			"groupEventsByKind": groupEventsByKind,
			"homeEventGroups":   homeEventGroups,
			"orgsBySocial":      orgsBySocial,
		}

		for _, p := range pages {
			tmpl, err := template.New(filepath.Base(p.tmplFiles[0])).Funcs(funcMap).ParseFiles(p.tmplFiles...)
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
