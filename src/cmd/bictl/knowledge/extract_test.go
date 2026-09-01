package knowledge

import "testing"

func TestURLToFilename(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{"https://example.org", "example-org.txt"},
		{"https://example.org/about/us", "example-org-about-us.txt"},
		{"https://example.org/", "example-org.txt"},
		{"%%invalid", "unknown.txt"},
	}
	for _, tt := range tests {
		got := URLToFilename(tt.url)
		if got != tt.want {
			t.Errorf("URLToFilename(%q) = %q, want %q", tt.url, got, tt.want)
		}
	}
}
