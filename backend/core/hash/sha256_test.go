package hash

import "testing"

func TestSHA256(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"Hello, World!", "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"},
	}

	for _, tt := range tests {
		result := SHA256(tt.input)
		if result != tt.expected {
			t.Errorf("SHA256(%q) = %s, want %s", tt.input, result, tt.expected)
		}
	}
}
