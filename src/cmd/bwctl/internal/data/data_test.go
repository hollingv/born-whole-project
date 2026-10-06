package data

import (
	"reflect"
	"testing"
)

func TestOrganizationStructHas8Fields(t *testing.T) {
	const expected = 8
	got := reflect.TypeOf(Organization{}).NumField()
	if got != expected {
		t.Errorf("Organization struct has %d fields, expected %d", got, expected)
	}
}

func TestOrgIDReferencesAreValid(t *testing.T) {
	// Verify every event references a valid organization ID
	for _, event := range Events {
		if event.OrgID != "" && FindOrgByID(event.OrgID) == nil {
			t.Errorf("event %q references unknown OrgID %q", event.Description, event.OrgID)
		}
	}
	// Verify every news item references a valid organization ID
	for _, item := range NewsItems {
		if item.OrgID != "" && FindOrgByID(item.OrgID) == nil {
			t.Errorf("news item %q references unknown OrgID %q", item.Description, item.OrgID)
		}
	}
	// Verify every donation item references a valid organization ID
	for _, item := range DonationItems {
		if item.OrgID != "" && FindOrgByID(item.OrgID) == nil {
			t.Errorf("donation item %q references unknown OrgID %q", item.Description, item.OrgID)
		}
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
