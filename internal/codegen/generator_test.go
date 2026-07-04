package codegen

import (
	"context"
	"testing"
)

func TestSequentialGeneratorStartsAtMinimumLength(t *testing.T) {
	generator := NewSequentialGenerator()

	code, err := generator.Generate(context.Background())
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if len(code) != MinimumGeneratedCodeLength {
		t.Fatalf("Generate() length = %d, want %d; code=%q", len(code), MinimumGeneratedCodeLength, code)
	}
}

func TestSequentialGeneratorGeneratesUniqueCodes(t *testing.T) {
	generator := NewSequentialGenerator()

	first, err := generator.Generate(context.Background())
	if err != nil {
		t.Fatalf("Generate() first error = %v", err)
	}

	second, err := generator.Generate(context.Background())
	if err != nil {
		t.Fatalf("Generate() second error = %v", err)
	}

	if first == second {
		t.Fatalf("Generate() returned duplicate codes: %q", first)
	}
}
