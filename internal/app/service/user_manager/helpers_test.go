package usermanager

import (
	"reflect"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/app/service/user_manager/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type fixture struct {
	ctx context.Context

	usersFacade    *mocks.UsersFacade
	implementation *Implementation
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		usersFacade: mocks.NewUsersFacade(t),
	}

	fx.implementation = &Implementation{
		usersFacade: fx.usersFacade,
	}

	return fx
}

func Test_convertModelUserToProtoUser(t *testing.T) {
	pbSex := usermanagerpb.Sex_SEX_MALE
	type args struct {
		user model.User
	}
	tests := []struct {
		name string
		args args
		want *usermanagerpb.User
	}{
		{
			name: "tc1",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "name",
					TelegramID: 111,
					Email:      maybe.Just("email"),
					Phone:      maybe.Just("phone"),
					State:      model.UserStateRegistered,
					Birthdate:  maybe.Just("1990-01-30"),
					Sex:        maybe.Just(model.SexMale),
				},
			},
			want: &usermanagerpb.User{
				Id:         1,
				Name:       "name",
				TelegramId: 111,
				Email:      wrapperspb.String("email"),
				Phone:      wrapperspb.String("phone"),
				State:      usermanagerpb.UserState_USER_STATE_REGISTERED,
				Birthdate:  wrapperspb.String("1990-01-30"),
				Sex:        &pbSex,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelUserToProtoUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelUserToProtoUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertProtoUserToModelUser(t *testing.T) {
	pbSex := usermanagerpb.Sex_SEX_MALE
	type args struct {
		user *usermanagerpb.User
	}
	tests := []struct {
		name string
		args args
		want model.User
	}{
		{
			name: "tc1",
			args: args{
				user: &usermanagerpb.User{
					Id:         1,
					Name:       "name",
					TelegramId: 100,
					Email:      wrapperspb.String("email"),
					Phone:      wrapperspb.String("phone"),
					State:      usermanagerpb.UserState_USER_STATE_CHANGING_EMAIL,
					Birthdate:  wrapperspb.String("1990-01-30"),
					Sex:        &pbSex,
				},
			},
			want: model.User{
				ID:         1,
				Name:       "name",
				TelegramID: 100,
				Email:      maybe.Just("email"),
				Phone:      maybe.Just("phone"),
				State:      model.UserStateChangingEmail,
				Birthdate:  maybe.Just("1990-01-30"),
				Sex:        maybe.Just(model.SexMale),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProtoUserToModelUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProtoUserToModelUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getErrorDetails(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		args args
		want *errorDetails
	}{
		{
			name: "keys is nil",
			args: args{
				keys: nil,
			},
			want: nil,
		},
		{
			name: "keys is empty",
			args: args{
				keys: []string{},
			},
			want: nil,
		},
		{
			name: "Name",
			args: args{
				keys: []string{"Name"},
			},
			want: &errorDetails{
				Reason: reasonInvalidUserName,
				Lexeme: invalidUserNameLexeme,
			},
		},
		{
			name: "TelegramID",
			args: args{
				keys: []string{"TelegramID"},
			},
			want: &errorDetails{
				Reason: reasonInvalidUserTelegramID,
				Lexeme: invalidUserTelegramIDLexeme,
			},
		},
		{
			name: "Email",
			args: args{
				keys: []string{"Email"},
			},
			want: &errorDetails{
				Reason: reasonInvalidUserEmail,
				Lexeme: invalidUserEmailLexeme,
			},
		},
		{
			name: "Phone",
			args: args{
				keys: []string{"Phone"},
			},
			want: &errorDetails{
				Reason: reasonInvalidUserPhone,
				Lexeme: invalidUserPhoneLexeme,
			},
		},
		{
			name: "State",
			args: args{
				keys: []string{"State"},
			},
			want: &errorDetails{
				Reason: reasonInvalidUserState,
				Lexeme: invalidUserStateLexeme,
			},
		},
		{
			name: "Birthdate",
			args: args{
				keys: []string{"Birthdate"},
			},
			want: &errorDetails{
				Reason: reasonInvalidUserBirthdate,
				Lexeme: invalidUserBirthdateLexeme,
			},
		},
		{
			name: "Sex",
			args: args{
				keys: []string{"Sex"},
			},
			want: &errorDetails{
				Reason: reasonInvalidUserSex,
				Lexeme: invalidUserSexLexeme,
			},
		},
		{
			name: "not found",
			args: args{
				keys: []string{"not found"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getErrorDetails(tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getErrorDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateBirthdate(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "not MaybeString",
			args: args{
				value: "",
			},
			wantErr: true,
		},
		{
			name: "not valid and empty",
			args: args{
				value: maybe.Nothing[string](),
			},
			wantErr: false,
		},
		{
			name: "valid and empty",
			args: args{
				value: maybe.Just(""),
			},
			wantErr: true,
		},
		{
			name: "valid and has invalid value",
			args: args{
				value: maybe.Just("invalid value"),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: maybe.Just("1990-01-30"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateBirthdate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateBirthdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateEmail(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "not MaybeString",
			args: args{
				value: "",
			},
			wantErr: true,
		},
		{
			name: "not valid and empty",
			args: args{
				value: maybe.Nothing[string](),
			},
			wantErr: false,
		},
		{
			name: "valid and empty",
			args: args{
				value: maybe.Just(""),
			},
			wantErr: true,
		},
		{
			name: "valid and has invalid value",
			args: args{
				value: maybe.Just("invalid value"),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: maybe.Just("email@email.ru"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateEmail(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatePhone(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "not MaybeString",
			args: args{
				value: "",
			},
			wantErr: true,
		},
		{
			name: "not valid and empty",
			args: args{
				value: maybe.Nothing[string](),
			},
			wantErr: false,
		},
		{
			name: "valid and empty",
			args: args{
				value: maybe.Just(""),
			},
			wantErr: true,
		},
		{
			name: "valid and has invalid value",
			args: args{
				value: maybe.Just("invalid value"),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: maybe.Just("+79998887766"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePhone(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validatePhone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateUserSex(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "not Maybe[int32]",
			args: args{
				value: "",
			},
			wantErr: true,
		},
		{
			name: "not valid and empty",
			args: args{
				value: maybe.Nothing[model.Sex](),
			},
			wantErr: false,
		},
		{
			name: "valid and empty",
			args: args{
				value: maybe.Just(model.Sex(0)),
			},
			wantErr: true,
		},
		{
			name: "valid and eq 0",
			args: args{
				value: maybe.Just(model.Sex(0)),
			},
			wantErr: true,
		},
		{
			name: "gt numberOfSexes",
			args: args{
				value: maybe.Just(model.Sex(100)),
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				value: maybe.Just(model.SexMale),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateUserSex(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateUserSex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
