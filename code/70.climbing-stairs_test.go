package code

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "Example 4",
			args: args{
				n: 4,
			},
			want: 5,
		},
		{
			name: "Example 5",
			args: args{
				n: 45,
			},
			want: 1836311903,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
