package goevil

import "testing"

func TestLevenshtein(t *testing.T) {
	type args struct {
	}
	tests := map[string]struct {
		str1 string
		str2 string
		want int
	}{
		"same": {
			str1: "batman",
			str2: "batman",
			want: 0,
		},
		"diff": {
			str1: "superman",
			str2: "superboy",
			want: 3,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := distance(tt.str1, tt.str2); got != tt.want {
				t.Errorf("levenshtein() = %v, want %v", got, tt.want)
			}
		})
	}
}
