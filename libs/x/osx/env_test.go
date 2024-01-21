package osx

import (
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
	// Test case: Environment variable exists
	const testKey1 = "TEST_ENV_VAR"
	const testValue1 = "test_value"
	os.Setenv(testKey1, testValue1)

	if got := Getenv(testKey1, "default"); got != testValue1 {
		t.Errorf("Getenv(%q, \"default\") = %q, want %q", testKey1, got, testValue1)
	}

	// Test case: Environment variable does not exist, return default
	const testKey2 = "NON_EXISTENT_ENV_VAR"
	const defaultValue = "default_value"

	if got := Getenv(testKey2, defaultValue); got != defaultValue {
		t.Errorf("Getenv(%q, %q) = %q, want %q", testKey2, defaultValue, got, defaultValue)
	}

	// Test case: Environment variable exists but is empty, return default
	const testKey3 = "EMPTY_ENV_VAR"
	os.Setenv(testKey3, "")

	if got := Getenv(testKey3, defaultValue); got != defaultValue {
		t.Errorf("Getenv(%q, %q) = %q, want %q", testKey3, defaultValue, got, defaultValue)
	}
}
