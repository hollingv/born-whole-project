package main

import (
	"testing"

	"iwebsite/src/cmd/bictl/internal/data"
)

func TestEachGroupHasOrganizations(t *testing.T) {
	for _, group := range data.OrgGroups {
		if len(group.Organizations) == 0 {
			t.Errorf("group %q has no organizations", group.Type)
		}
	}
}
