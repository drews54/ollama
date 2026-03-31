package imagegen

import "testing"

func TestImageGenAllowOversubscribe(t *testing.T) {
	tests := []struct {
		name   string
		value  *string
		expect bool
	}{
		{
			name:   "unset",
			value:  nil,
			expect: false,
		},
		{
			name:   "one",
			value:  ptr("1"),
			expect: true,
		},
		{
			name:   "true",
			value:  ptr("true"),
			expect: true,
		},
		{
			name:   "zero",
			value:  ptr("0"),
			expect: false,
		},
		{
			name:   "invalid",
			value:  ptr("banana"),
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value == nil {
				t.Setenv("OLLAMA_IMAGEGEN_ALLOW_OVERSUBSCRIBE", "")
			} else {
				t.Setenv("OLLAMA_IMAGEGEN_ALLOW_OVERSUBSCRIBE", *tt.value)
			}

			got := imageGenAllowOversubscribe()
			if got != tt.expect {
				t.Fatalf("imageGenAllowOversubscribe() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func ptr(s string) *string {
	return &s
}
