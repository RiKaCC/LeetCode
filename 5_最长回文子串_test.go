package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"1",
			args:args{s:"acbcakk"},
			want:"acbca",
		},
		{
			name:"2",
			args:args{s:"ccccc"},
			want:"ccccc",
		},
		{
			name:"3",
			args:args{s:"asfasf"},
			want:"a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

