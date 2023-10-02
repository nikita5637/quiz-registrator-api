package model

import (
	"reflect"
	"testing"
	"time"

	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
)

func TestDateTime_AsTime(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	tests := []struct {
		name string
		d    DateTime
		want time.Time
	}{
		{
			name: "tc1",
			d:    DateTime(timeutils.TimeNow()),
			want: timeutils.TimeNow(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AsTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.AsTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_String(t *testing.T) {
	tests := []struct {
		name string
		d    DateTime
		want string
	}{
		{
			name: "tc1",
			d:    DateTime(time.Unix(1, 0).UTC()),
			want: "1970-01-01 00:00:01 +0000 UTC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_MarshalJSON(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	tests := []struct {
		name    string
		d       DateTime
		want    []byte
		wantErr bool
	}{
		{
			name:    "tc1",
			d:       DateTime(timeutils.TimeNow()),
			want:    []byte("\"2006-01-02T15:04:00Z\""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_UnmarshalJSON(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	dateTime := DateTime(timeutils.TimeNow())

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		d       *DateTime
		args    args
		wantErr bool
	}{
		{
			name: "error",
			d:    &dateTime,
			args: args{
				data: []byte("invalid"),
			},
			wantErr: true,
		},
		{
			name: "ok",
			d:    &dateTime,
			args: args{
				data: []byte("\"2006-01-02T15:04:00Z\""),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateDateTime(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "is not DateTime",
			args: args{
				value: "is not DateTime",
			},
			wantErr: true,
		},
		{
			name: "is empty",
			args: args{
				value: DateTime(time.Unix(0, 0).UTC()),
			},
			wantErr: true,
		},
		{
			name: "eq -1",
			args: args{
				value: DateTime(time.Unix(-1, 0)),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: DateTime(timeutils.TimeNow()),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateDateTime(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateDateTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
