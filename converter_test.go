package belka

import "testing"

func TestSm(t *testing.T) {
	type testCase struct {
		value   string
		want    int32
		wantErr bool
	}
	tests := []testCase{
		{"123", 12300, false},
		{"abc", 0, true},
		{"123.00", 12300, false},
		{"22.1", 2210, false},
		{"6542.11", 654211, false},
		{"6542.119", 654211, false},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got, err := Sm[string, int32](tt.value)

			if got != tt.want {
				t.Errorf("Sm() got = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Sm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_multipier(t *testing.T) {
	type testCase struct {
		name  string
		value rune
		want  rune
	}
	tests := []testCase{
		{"2", 2, 100},
		{"1", 1, 10},
		{"0", 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplier[rune](tt.value); got != tt.want {
				t.Errorf("multiplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
