package imagegen

import "testing"

func TestValidateImageRunnerMemoryFit(t *testing.T) {
	t.Run("strict disabled allows oversubscribe", func(t *testing.T) {
		err := validateImageRunnerMemoryFit(32<<30, 30<<30, false)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})

	t.Run("strict enabled rejects when under required memory", func(t *testing.T) {
		err := validateImageRunnerMemoryFit(32<<30, 30<<30, true)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("strict enabled allows when unknown available memory", func(t *testing.T) {
		err := validateImageRunnerMemoryFit(32<<30, 0, true)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
}
