package main

import "testing"

func TestClientValidFlags(t *testing.T) {
	// should be invalid
	invalidFlags := clientFlagsValid(true, true, true)
	if invalidFlags {
		t.Fatalf("Too many true flags should fail!")
	}

	// should be valid
	validFlags := clientFlagsValid(false, false, true)
	if !validFlags {
		t.Fatalf("One flag should be set as valid!")
	}
}
