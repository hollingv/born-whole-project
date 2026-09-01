package youtube

import "testing"

func TestChannelConfigured(t *testing.T) {
	if Channel == "" {
		t.Error("Channel handle must not be empty")
	}
}
