package code

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "Hello World",
			},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{
				s: "   fly me   to   the moon  ",
			},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{
				s: "luffy is still joyboy",
			},
			want: 6,
		},
		{
			name: "Example 4",
			args: args{
				s: "            as ",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord(%s) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
