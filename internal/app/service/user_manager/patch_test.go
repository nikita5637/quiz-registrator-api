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
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_PatchUser(t *testing.T) {
	t.Run("get original user internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{}, errors.New("some error"))

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:   1,
				Name: "name",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("original user not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{}, users.ErrUserNotFound)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:   1,
				Name: "name",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonUserNotFound, errorInfo.Reason)
	})

	t.Run("validation error. invalid user name alphabet", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      maybe.Nothing[string](),
			Phone:      maybe.Nothing[string](),
			State:      model.UserStateChangingName,
			Birthdate:  maybe.Nothing[string](),
			Sex:        maybe.Nothing[model.Sex](),
		}, nil)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:   1,
				Name: "name",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
				},
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

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      maybe.Nothing[string](),
			Phone:      maybe.Nothing[string](),
			State:      model.UserStateChangingName,
			Birthdate:  maybe.Nothing[string](),
			Sex:        maybe.Nothing[model.Sex](),
		}, nil)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:   1,
				Name: "абвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгда",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
				},
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

	t.Run("validation error. invalid telegram id", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      maybe.Nothing[string](),
			Phone:      maybe.Nothing[string](),
			State:      model.UserStateChangingName,
			Birthdate:  maybe.Nothing[string](),
			Sex:        maybe.Nothing[model.Sex](),
		}, nil)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: 0,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
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

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			Email:      maybe.Nothing[string](),
			Phone:      maybe.Nothing[string](),
			TelegramID: -100,
			State:      model.UserStateChangingName,
			Birthdate:  maybe.Nothing[string](),
			Sex:        maybe.Nothing[model.Sex](),
		}, nil)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      wrapperspb.String("invalid email"),
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"email",
				},
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

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      maybe.Nothing[string](),
			Phone:      maybe.Nothing[string](),
			State:      model.UserStateChangingName,
			Birthdate:  maybe.Nothing[string](),
			Sex:        maybe.Nothing[model.Sex](),
		}, nil)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      wrapperspb.String("email@example.com"),
				Phone:      wrapperspb.String("invalid phone"),
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"phone",
				},
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

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      maybe.Nothing[string](),
			Phone:      maybe.Nothing[string](),
			State:      model.UserStateChangingName,
			Birthdate:  maybe.Nothing[string](),
			Sex:        maybe.Nothing[model.Sex](),
		}, nil)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      wrapperspb.String("email@example.com"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState_USER_STATE_INVALID,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"state",
				},
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

	t.Run("validation error. invalid birthdate", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      maybe.Nothing[string](),
			Phone:      maybe.Nothing[string](),
			State:      model.UserStateChangingName,
			Birthdate:  maybe.Nothing[string](),
			Sex:        maybe.Nothing[model.Sex](),
		}, nil)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      wrapperspb.String("email@example.com"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
				Birthdate:  wrapperspb.String("invalid value"),
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"birthdate",
				},
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

	t.Run("validation error. invalid sex", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_INVALID

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			Email:      maybe.Nothing[string](),
			Phone:      maybe.Nothing[string](),
			TelegramID: -100,
			State:      model.UserStateChangingName,
			Birthdate:  maybe.Nothing[string](),
			Sex:        maybe.Nothing[model.Sex](),
		}, nil)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      wrapperspb.String("email@example.com"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
				Birthdate:  wrapperspb.String("1990-01-30"),
				Sex:        &pbSex,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"sex",
				},
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

	t.Run("internal error while patch user", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_MALE

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			State:      model.UserStateChangingName,
		}, nil)

		fx.usersFacade.EXPECT().PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      maybe.Just("email@example.com"),
			Phone:      maybe.Just("+79998887766"),
			State:      model.UserStateWelcome,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexMale),
		}).Return(model.User{}, errors.New("some error"))

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      wrapperspb.String("email@example.com"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
				Birthdate:  wrapperspb.String("1990-01-30"),
				Sex:        &pbSex,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
					"email",
					"phone",
					"state",
					"birthdate",
					"sex",
				},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("user not found error while patch user", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_FEMALE

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			State:      model.UserStateChangingName,
		}, nil)

		fx.usersFacade.EXPECT().PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      maybe.Just("email@example.com"),
			Phone:      maybe.Just("+79998887766"),
			State:      model.UserStateWelcome,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexFemale),
		}).Return(model.User{}, users.ErrUserTelegramIDAlreadyExists)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      wrapperspb.String("email@example.com"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
				Birthdate:  wrapperspb.String("1990-01-30"),
				Sex:        &pbSex,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
					"email",
					"phone",
					"state",
					"birthdate",
					"sex",
				},
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

	t.Run("user with specified email error while patch user", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_FEMALE

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			State:      model.UserStateChangingName,
		}, nil)

		fx.usersFacade.EXPECT().PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      maybe.Just("email@example.com"),
			Phone:      maybe.Just("+79998887766"),
			State:      model.UserStateWelcome,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexFemale),
		}).Return(model.User{}, users.ErrUserEmailAlreadyExists)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      wrapperspb.String("email@example.com"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
				Birthdate:  wrapperspb.String("1990-01-30"),
				Sex:        &pbSex,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
					"email",
					"phone",
					"state",
					"birthdate",
					"sex",
				},
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

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			State:      model.UserStateChangingName,
		}, nil)

		fx.usersFacade.EXPECT().PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      maybe.Just("email@example.com"),
			Phone:      maybe.Just("+79998887766"),
			State:      model.UserStateWelcome,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexFemale),
		}).Return(model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      maybe.Just("email@example.com"),
			Phone:      maybe.Just("+79998887766"),
			State:      model.UserStateWelcome,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexFemale),
		}, nil)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      wrapperspb.String("email@example.com"),
				Phone:      wrapperspb.String("+79998887766"),
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
				Birthdate:  wrapperspb.String("1990-01-30"),
				Sex:        &pbSex,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
					"email",
					"phone",
					"state",
					"birthdate",
					"sex",
				},
			},
		})
		assert.Equal(t, &usermanagerpb.User{
			Id:         1,
			Name:       "Имя",
			TelegramId: -100,
			Email:      wrapperspb.String("email@example.com"),
			Phone:      wrapperspb.String("+79998887766"),
			State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			Birthdate:  wrapperspb.String("1990-01-30"),
			Sex:        &pbSex,
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validatePatchedUser(t *testing.T) {
	type args struct {
		user                model.User
		onlyRussianAlphabet bool
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
					Name:       "Name",
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
			name: "empty name",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "",
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					TelegramID: 100,
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
					ID:         1,
					Name:       "abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdea",
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					TelegramID: 100,
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
					ID:         1,
					Name:       "абвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгда",
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					TelegramID: 100,
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "telegram ID eq 0",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Name",
					TelegramID: 0,
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
			name: "state eq 0",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      0,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "state gt numberOfUserStates",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      100,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "email is not valid",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Name",
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
					ID:         1,
					Name:       "Name",
					TelegramID: 100,
					Email:      maybe.Just(""),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "email is valid and has invalid value",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Name",
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
					ID:         1,
					Name:       "Name",
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
					ID:         1,
					Name:       "Name",
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
					ID:         1,
					Name:       "Name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Just(""),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "phone is valid and has invalid value",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Name",
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
					ID:         1,
					Name:       "Name",
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
					ID:         1,
					Name:       "Name",
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
					ID:         1,
					Name:       "Name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Just(""),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "birthdate is valid and has invalid value",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Just("invalid birthdate"),
					Sex:        maybe.Nothing[model.Sex](),
				},
			},
			wantErr: true,
		},
		{
			name: "birthdate is valid and has valid value",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Name",
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
					ID:         1,
					Name:       "Name",
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
					ID:         1,
					Name:       "Name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Just(model.Sex(0)),
				},
			},
			wantErr: true,
		},
		{
			name: "sex is male",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Name",
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
					ID:         1,
					Name:       "Name",
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
		{
			name: "error. english name",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Name",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Just(model.SexFemale),
				},
				onlyRussianAlphabet: true,
			},
			wantErr: true,
		},
		{
			name: "ok. english name",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Name",
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
		{
			name: "ok. russian name",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "Имя",
					TelegramID: 100,
					Email:      maybe.Nothing[string](),
					Phone:      maybe.Nothing[string](),
					State:      model.UserStateWelcome,
					Birthdate:  maybe.Nothing[string](),
					Sex:        maybe.Just(model.SexFemale),
				},
				onlyRussianAlphabet: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePatchedUser(tt.args.user, tt.args.onlyRussianAlphabet); (err != nil) != tt.wantErr {
				t.Errorf("validatePatchedUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
