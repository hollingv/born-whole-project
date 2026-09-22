package data

import (
	"reflect"
	"testing"
)

func TestOrganizationStructHas8Fields(t *testing.T) {
	const expected = 7
	got := reflect.TypeOf(Organization{}).NumField()
	if got != expected {
		t.Errorf("Organization struct has %d fields, expected %d", got, expected)
	}
}

func TestOrgGroupsNotEmpty(t *testing.T) {
	for _, g := range OrgGroups {
		for _, org := range g.Organizations {
			if org.Name == "" {
				t.Errorf("organization in group %q has empty Name", g.Type)
			}
			if org.Website == "" {
				t.Errorf("organization %q has empty Website", org.Name)
			}
		}
	}
}
