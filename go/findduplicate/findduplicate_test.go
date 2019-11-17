package findduplicate

import (
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}

	argss := args{[]int{3, 2, 1, 1}}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"name1", argss, 1},
		// {"name2", argss, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
