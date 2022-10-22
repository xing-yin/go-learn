package clause

import (
	"testing"
)

func Test_genBindVars(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				num: 3,
			},
			want: "?,?,?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genBindVars(tt.args.num); got != tt.want {
				t.Errorf("genBindVars() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestCastFields(t *testing.T) {
//	values := interface{}
//	strings.Join(values.([]string), ",")
//}
