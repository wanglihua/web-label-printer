package hash

import "testing"

func TestMd5(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"Hello, World!", "65a8e27d8879283831b664bd8b7f0ad4"},
	}

	for _, tt := range tests {
		result := Md5(tt.input)
		if result != tt.expected {
			t.Errorf("Md5(%q) = %s, want %s", tt.input, result, tt.expected)
		}
	}
}
