package domain

import "testing"

func TestNormalizeLongURL(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    string
		wantErr bool
	}{
		{name: "https", raw: "https://example.com/path", want: "https://example.com/path"},
		{name: "http", raw: "http://example.com", want: "http://example.com"},
		{name: "trims spaces", raw: "  https://example.com  ", want: "https://example.com"},
		{name: "missing scheme", raw: "example.com", wantErr: true},
		{name: "unsupported scheme", raw: "ftp://example.com", wantErr: true},
		{name: "missing host", raw: "https:///path", wantErr: true},
		{name: "empty", raw: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeLongURL(tt.raw)
			if tt.wantErr {
				if err == nil {
					t.Fatal("NormalizeLongURL() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("NormalizeLongURL() error = %v, want nil", err)
			}
			if got != tt.want {
				t.Fatalf("NormalizeLongURL() = %q, want %q", got, tt.want)
			}
		})
	}
}
