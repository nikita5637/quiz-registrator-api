package model

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestNewMaybeInt32(t *testing.T) {
	type args struct {
		value int32
	}
	tests := []struct {
		name string
		args args
		want MaybeInt32
	}{
		{
			name: "tc1",
			args: args{
				value: 1,
			},
			want: MaybeInt32{
				Valid: true,
				Value: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMaybeInt32(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMaybeInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaybeInt32_ToSQL(t *testing.T) {
	type fields struct {
		Valid bool
		Value int32
	}
	tests := []struct {
		name   string
		fields fields
		want   sql.NullInt32
	}{
		{
			name: "tc1",
			fields: fields{
				Valid: true,
				Value: 1,
			},
			want: sql.NullInt32{
				Int32: 1,
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := MaybeInt32{
				Valid: tt.fields.Valid,
				Value: tt.fields.Value,
			}
			if got := v.ToSQL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaybeInt32.ToSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaybeInt32_ToSQLNullInt64(t *testing.T) {
	type fields struct {
		Valid bool
		Value int32
	}
	tests := []struct {
		name   string
		fields fields
		want   sql.NullInt64
	}{
		{
			name: "tc1",
			fields: fields{
				Valid: true,
				Value: 1,
			},
			want: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := MaybeInt32{
				Valid: tt.fields.Valid,
				Value: tt.fields.Value,
			}
			if got := v.ToSQLNullInt64(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaybeInt32.ToSQLNullInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMaybeString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want MaybeString
	}{
		{
			name: "tc1",
			args: args{
				value: "value",
			},
			want: MaybeString{
				Valid: true,
				Value: "value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMaybeString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMaybeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaybeString_ToSQL(t *testing.T) {
	type fields struct {
		Valid bool
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   sql.NullString
	}{
		{
			name: "tc1",
			fields: fields{
				Valid: true,
				Value: "value",
			},
			want: sql.NullString{
				String: "value",
				Valid:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := MaybeString{
				Valid: tt.fields.Valid,
				Value: tt.fields.Value,
			}
			if got := v.ToSQL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaybeString.ToSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
