package iteration

import "testing"

func TestRepeat(t *testing.T) {
	tests := []struct {
		char   string
		count  int
		expect string
	}{
		{"a", 5, "aaaaa"},
		{"bb", 3, "bbbbbb"},
	}
	for _, tt := range tests {
		repeated := Repeat(tt.char, tt.count)
		if repeated != tt.expect {
			t.Errorf("expected %q but got %q for char '%s' and count %d",
				tt.expect, repeated, tt.char, tt.count)
		}
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
