package normalizephone

import "testing"

func TestNormalize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// TODO: Add test cases.
		{"6289898989", "6289898989", "6289898989"},
		{"08188772266", "08188772266", "628188772266"},
		{"+6289898989", "+6289898989", "6289898989"},
		{"1234", "1234", "-"},
		{"1", "1", "-"},
		{"", "", "-"},
		{"ababababababababa", "ababababababababa", "-"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Normalize(tt.input); got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
