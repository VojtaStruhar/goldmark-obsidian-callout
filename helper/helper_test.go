package helper

import "testing"

// The default callout type is Note.
// This is determined by Note being the first element in the CalloutType definition (iota)
func TestDefaultCalloutType(t *testing.T) {
	nonexistentCalloutTag := "non-existent-callout-type"
	calloutType := CalloutTypeMapping[nonexistentCalloutTag]
	if calloutType != Note {
		t.Fail()
	}
}
