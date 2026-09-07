package main

import (
	"testing"

	"born-whole-project/src/cmd/bwctl/internal/data"
)

func TestEachGroupHasOrganizations(t *testing.T) {
	for _, group := range data.OrgGroups {
		if len(group.Organizations) == 0 {
			t.Errorf("group %q has no organizations", group.Type)
		}
	}
}
