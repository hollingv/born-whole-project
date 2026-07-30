package main

import (
	"testing"
)

func TestEachGroupHasOrganizations(t *testing.T) {
	for _, group := range orgGroups {
		if len(group.Organizations) == 0 {
			t.Errorf("group %q has no organizations", group.Type)
		}
	}
}
