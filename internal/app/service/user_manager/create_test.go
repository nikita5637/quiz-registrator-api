package usermanager

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	users "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_CreateUser(t *testing.T) {
	t.Run("validation error. empty user name", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "",
				TelegramId: -100,
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidUserName, errorInfo.Reason)
	})

	t.Run("validation error. user name length gt 100", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghija",
				TelegramId: -100,
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidUserName, errorInfo.Reason)
	})

	t.Run("validation error. empty Telegram ID", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: 0,
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidUserTelegramID, errorInfo.Reason)
	})

	t.Run("validation error. invalid email format", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      wrapperspb.String("invalid email"),
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidUserEmail, errorInfo.Reason)
	})

	t.Run("validation error. invalid phone format", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      wrapperspb.String("email@email.ru"),
				Phone:      wrapperspb.String("invalid phone"),
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidUserPhone, errorInfo.Reason)
	})

	t.Run("validation error. invalid user state", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      wrapperspb.String("email@email.ru"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState_USER_STATE_INVALID,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidUserState, errorInfo.Reason)
	})

	t.Run("validation error. invalid user birthdate", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      wrapperspb.String("email@email.ru"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
				Birthdate:  wrapperspb.String("invalid birthdate"),
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidUserBirthdate, errorInfo.Reason)
	})

	t.Run("validation error. invalid user sex", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_INVALID

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      wrapperspb.String("email@email.ru"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
				Sex:        &pbSex,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidUserSex, errorInfo.Reason)
	})

	t.Run("internal error while create user", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_MALE

		fx.usersFacade.EXPECT().CreateUser(fx.ctx, model.User{
			Name:       "name",
			TelegramID: -100,
			Phone:      maybe.Just("+79998887766"),
			Email:      maybe.Just("email@email.ru"),
			State:      model.UserStateWelcome,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexMale),
		}).Return(model.User{}, errors.New("some error"))

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      wrapperspb.String("email@email.ru"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState(model.UserStateWelcome),
				Birthdate:  wrapperspb.String("1990-01-30"),
				Sex:        &pbSex,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("user with specified Telegram ID already exists error while create user", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_FEMALE

		fx.usersFacade.EXPECT().CreateUser(fx.ctx, model.User{
			Name:       "name",
			TelegramID: -100,
			Phone:      maybe.Just("+79998887766"),
			Email:      maybe.Just("email@email.ru"),
			State:      model.UserStateWelcome,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexFemale),
		}).Return(model.User{}, users.ErrUserTelegramIDAlreadyExists)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      wrapperspb.String("email@email.ru"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState(model.UserStateWelcome),
				Birthdate:  wrapperspb.String("1990-01-30"),
				Sex:        &pbSex,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonUserAlreadyExists, errorInfo.Reason)
	})

	t.Run("user with specified email already exists error while create user", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_FEMALE

		fx.usersFacade.EXPECT().CreateUser(fx.ctx, model.User{
			Name:       "name",
			TelegramID: -100,
			Phone:      maybe.Just("+79998887766"),
			Email:      maybe.Just("email@email.ru"),
			State:      model.UserStateWelcome,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexFemale),
		}).Return(model.User{}, users.ErrUserEmailAlreadyExists)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      wrapperspb.String("email@email.ru"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState(model.UserStateWelcome),
				Birthdate:  wrapperspb.String("1990-01-30"),
				Sex:        &pbSex,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonUserAlreadyExists, errorInfo.Reason)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_FEMALE

		fx.usersFacade.EXPECT().CreateUser(fx.ctx, model.User{
			Name:       "name",
			TelegramID: -100,
			Phone:      maybe.Just("+79998887766"),
			Email:      maybe.Just("email@email.ru"),
			State:      model.UserStateWelcome,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexFemale),
		}).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Phone:      maybe.Just("+79998887766"),
			Email:      maybe.Just("email@email.ru"),
			State:      model.UserStateWelcome,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexFemale),
		}, nil)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      wrapperspb.String("email@email.ru"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState(model.UserStateWelcome),
				Birthdate:  wrapperspb.String("1990-01-30"),
				Sex:        &pbSex,
			},
		})
		assert.NotNil(t, got)
		assert.Equal(t, &usermanagerpb.User{
			Id:         1,
			Name:       "name",
			TelegramId: -100,
			Email:      wrapperspb.String("email@email.ru"),
			Phone:      wrapperspb.String("+79998887766"),
			State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			Birthdate:  wrapperspb.String("1990-01-30"),
			Sex:        &pbSex,
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validateCreatedUser(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no ID",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: false,
		},
		{
			name: "empty name",
			args: args{
				user: model.User{
					Name:       "",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "name is too long",
			args: args{
				user: model.User{
					Name:       "abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdea",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "name is too long #2",
			args: args{
				user: model.User{
					Name:       "абвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгда",
					TelegramID: 100,
					State:      model.UserStateWelcome,
				},
			},
			wantErr: true,
		},
		{
			name: "telegram ID eq 0",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 0,
					State:      model.UserStateWelcome,
				},
			},
			wantErr: true,
		},
		{
			name: "state eq 0",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					State:      0,
				},
			},
			wantErr: true,
		},
		{
			name: "state gt numberOfUserStates",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					State:      100,
				},
			},
			wantErr: true,
		},
		{
			name: "email is not valid",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 1,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: false,
		},
		{
			name: "email is valid and empty",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Just(""),
					State:      model.UserStateWelcome,
				},
			},
			wantErr: true,
		},
		{
			name: "email is valid and has invalid value",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Just("invalid email"),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "email is valid and has valid value",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Just("email@mail.ru"),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: false,
		},
		{
			name: "phone is not valid",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 1,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: false,
		},
		{
			name: "phone is valid and empty",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Phone:      maybe.Just(""),
					State:      model.UserStateWelcome,
				},
			},
			wantErr: true,
		},
		{
			name: "phone is valid and has invalid value",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Just("invalid phone"),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "phone is valid and has valid value",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Just("+79998887766"),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: false,
		},
		{
			name: "birthdate is not valid",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 1,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: false,
		},
		{
			name: "birthdate is valid and empty",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Just(""),
				},
			},
			wantErr: true,
		},
		{
			name: "birthdate is valid and has invalid value",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Just("invalid birthdate"),
				},
			},
			wantErr: true,
		},
		{
			name: "birthdate is valid and has valid value",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Just("1990-12-30"),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: false,
		},
		{
			name: "sex is not valid",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: false,
		},
		{
			name: "sex eq 0",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					State:      model.UserStateWelcome,
					Sex:        maybe.Just(model.Sex(0)),
				},
			},
			wantErr: true,
		},
		{
			name: "sex is male",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Just(model.SexMale),
				},
			},
			wantErr: false,
		},
		{
			name: "sex is female",
			args: args{
				user: model.User{
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Just(model.SexFemale),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreatedUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("validateCreatedUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
