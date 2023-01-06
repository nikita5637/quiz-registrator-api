package mysql

import "testing"

func Test_getOrderStmt(t *testing.T) {
	type args struct {
		sort string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				sort: "",
			},
			want: "",
		},
		{
			name: "test case 2",
			args: args{
				sort: "field",
			},
			want: "ORDER BY field ASC ",
		},
		{
			name: "test case 3",
			args: args{
				sort: "field1,field2",
			},
			want: "ORDER BY field1 ASC , field2 ASC ",
		},
		{
			name: "test case 4",
			args: args{
				sort: "field1,-field2",
			},
			want: "ORDER BY field1 ASC , field2 DESC ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOrderStmt(tt.args.sort); got != tt.want {
				t.Errorf("getOrderStmt() = %v, want %v", got, tt.want)
			}
		})
	}
}
