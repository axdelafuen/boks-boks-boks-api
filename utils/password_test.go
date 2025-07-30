package utils

import (
	"strings"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	// Check that the hash is not empty
	if hashedPassword == "" {
		t.Error("Expected non-empty hashed password")
	}

	// Check that the hash is different from the original password
	if hashedPassword == password {
		t.Error("Expected hashed password to be different from original")
	}

	// Check that bcrypt hash starts with expected prefix
	if !strings.HasPrefix(hashedPassword, "$2a$") && !strings.HasPrefix(hashedPassword, "$2b$") {
		t.Error("Expected bcrypt hash to start with $2a$ or $2b$")
	}
}

func TestComparePassword(t *testing.T) {
	password := "testpassword123"
	wrongPassword := "wrongpassword"

	// Hash the correct password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	// Test correct password comparison
	err = ComparePassword(hashedPassword, password)
	if err != nil {
		t.Errorf("Expected password comparison to succeed, but got error: %v", err)
	}

	// Test wrong password comparison
	err = ComparePassword(hashedPassword, wrongPassword)
	if err == nil {
		t.Error("Expected password comparison to fail with wrong password")
	}
}

func TestHashPasswordConsistency(t *testing.T) {
	password := "testpassword123"

	// Hash the same password multiple times
	hash1, err1 := HashPassword(password)
	hash2, err2 := HashPassword(password)

	if err1 != nil || err2 != nil {
		t.Fatalf("HashPassword failed: %v, %v", err1, err2)
	}

	// Hashes should be different (due to salt)
	if hash1 == hash2 {
		t.Error("Expected different hashes for same password (salt should make them unique)")
	}

	// But both should verify against the original password
	if err := ComparePassword(hash1, password); err != nil {
		t.Errorf("First hash should verify: %v", err)
	}

	if err := ComparePassword(hash2, password); err != nil {
		t.Errorf("Second hash should verify: %v", err)
	}
}

func TestIsValidPasswordLength(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"short", false},
		{"1234567", false}, // 7 chars
		{"12345678", true}, // 8 chars (minimum)
		{"longerpassword", true},
		{"", false},
	}

	for _, test := range tests {
		result := IsValidPasswordLength(test.password)
		if result != test.expected {
			t.Errorf("IsValidPasswordLength(%q) = %v, expected %v", test.password, result, test.expected)
		}
	}
}
