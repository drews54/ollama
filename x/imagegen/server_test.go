package imagegen

import (
	"errors"
	"testing"

	"github.com/ollama/ollama/llm"
)

func TestValidateImageModelMemoryFit(t *testing.T) {
	t.Run("strict fit disabled allows oversubscribe", func(t *testing.T) {
		err := validateImageModelMemoryFit(10, 1, 0, false, false)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("strict fit enabled with requireFull returns ErrLoadRequiredFull", func(t *testing.T) {
		err := validateImageModelMemoryFit(10, 1, 0, true, true)
		if !errors.Is(err, llm.ErrLoadRequiredFull) {
			t.Fatalf("expected ErrLoadRequiredFull, got %v", err)
		}
	})

	t.Run("strict fit enabled returns descriptive error when model does not fit", func(t *testing.T) {
		err := validateImageModelMemoryFit(10, 1, 2, false, true)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("strict fit enabled allows fitting model", func(t *testing.T) {
		err := validateImageModelMemoryFit(10, 10, 0, false, true)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})
}
