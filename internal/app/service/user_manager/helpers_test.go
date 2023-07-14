package usermanager

import (
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/service/user_manager/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"golang.org/x/net/context"
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
					Email:      model.NewMaybeString("email"),
					Phone:      model.NewMaybeString("phone"),
					State:      model.UserStateRegistered,
				},
			},
			want: &usermanagerpb.User{
				Id:         1,
				Name:       "name",
				TelegramId: 111,
				Email:      "email",
				Phone:      "phone",
				State:      usermanagerpb.UserState_USER_STATE_REGISTERED,
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
