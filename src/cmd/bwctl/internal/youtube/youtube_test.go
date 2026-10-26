package youtube

import "testing"

func TestChannelConfigured(t *testing.T) {
	if Channel == "" {
		t.Error("Channel handle must not be empty")
	}
}

func TestStripNonASCII(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"Hello World", "Hello World"},
		{"Caf\u00e9", "Caf"},                                 // accented character removed
		{"Hello \U0001F600 World", "Hello  World"},           // emoji removed
		{"\u4e2d\u6587", ""},                                 // Chinese characters removed
		{"  trim me  ", "trim me"},                           // surrounding whitespace trimmed
		{"tab\there", "tabhere"},                             // tab (ASCII 9) removed
		{"normal & symbols: 100%", "normal & symbols: 100%"}, // valid punctuation kept
		{"", ""},                           // empty string
		{"it&#39;s a test", "it's a test"}, // HTML entity decoded
		{"&amp;more", "&more"},             // &amp; decoded
		{"&quot;quoted&quot;", `"quoted"`}, // &quot; decoded
		{"caf&#233; &#128514;", "caf"},     // entity decoded then non-ASCII stripped
	}
	for _, c := range cases {
		got := stripNonASCII(c.input)
		if got != c.expected {
			t.Errorf("stripNonASCII(%q) = %q, want %q", c.input, got, c.expected)
		}
	}
}
