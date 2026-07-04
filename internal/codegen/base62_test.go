package codegen

import "testing"

func TestEncodeBase62(t *testing.T) {
	tests := []struct {
		name string
		id   uint64
		want string
	}{
		{name: "zero", id: 0, want: "a"},
		{name: "last single character", id: 61, want: "9"},
		{name: "first two characters", id: 62, want: "ba"},
		{name: "next two characters", id: 63, want: "bb"},
		{name: "larger value", id: 3844, want: "baa"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EncodeBase62(tt.id)
			if got != tt.want {
				t.Fatalf("EncodeBase62(%d) = %q, want %q", tt.id, got, tt.want)
			}
		})
	}
}

func TestMinimumGeneratedID(t *testing.T) {
	got := MinimumGeneratedID()
	if got != 14776336 {
		t.Fatalf("MinimumGeneratedID() = %d, want 14776336", got)
	}

	code := EncodeBase62(got)
	if len(code) != MinimumGeneratedCodeLength {
		t.Fatalf("minimum generated code length = %d, want %d; code=%q", len(code), MinimumGeneratedCodeLength, code)
	}
}
