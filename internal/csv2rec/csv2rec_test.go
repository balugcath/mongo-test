package csv2rec

import (
	"reflect"
	"testing"

	"github.com/balugcath/mongo-test/pkg/types"
)

func TestArray2Records(t *testing.T) {
	type args struct {
		a [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    []types.ProductPrice
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{a: nil},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{a: [][]string{{"a", "1"}}},
			want:    []types.ProductPrice{{ProductName: "a", Price: 1}},
			wantErr: false,
		},
		{
			name:    "test3",
			args:    args{a: [][]string{{"a", "1"}, {"b", "2.2"}}},
			want:    []types.ProductPrice{{ProductName: "a", Price: 1}, {ProductName: "b", Price: 2.2}},
			wantErr: false,
		},
		{
			name:    "test4",
			args:    args{a: [][]string{{"a", "1.1.1"}}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Array2Records(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("Array2Records() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Array2Records() = %v, want %v", got, tt.want)
			}
		})
	}
}
