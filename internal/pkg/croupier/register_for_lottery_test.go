package croupier

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"

	"github.com/stretchr/testify/assert"
)

func TestCroupier_RegisterForLottery(t *testing.T) {
	t.Run("lottery not implemented", func(t *testing.T) {
		croupier := New(Config{})

		ctx := context.Background()

		got, err := croupier.RegisterForLottery(ctx, model.Game{}, model.User{})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.Equal(t, model.ErrLotteryNotImplemented, errors.Unwrap(err))
	})

	t.Run("lottery is not active", func(t *testing.T) {
		croupier := New(Config{})

		globalConfig := config.GlobalConfig{}
		globalConfig.LotteryStartsBefore = 3600
		config.UpdateGlobalConfig(globalConfig)

		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-08 13:39")
		}

		ctx := context.Background()

		got, err := croupier.RegisterForLottery(ctx, model.Game{
			Date:     model.DateTime(time_utils.ConvertTime("2022-01-09 16:30")),
			LeagueID: model.LeagueQuizPlease,
			My:       true,
		}, model.User{})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.Equal(t, model.ErrLotteryNotAvailable, errors.Unwrap(err))
	})

	t.Run("error while registration, quiz please", func(t *testing.T) {
		quizPleaseCroupierMock := mocks.NewLotteryRegistrator(t)
		croupier := New(Config{
			QuizPleaseCroupier: quizPleaseCroupierMock,
		})

		globalConfig := config.GlobalConfig{}
		globalConfig.LotteryStartsBefore = 3600
		config.UpdateGlobalConfig(globalConfig)

		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-10 15:31")
		}

		ctx := context.Background()

		game := model.Game{
			Date:     model.DateTime(time_utils.ConvertTime("2022-01-09 16:30")),
			LeagueID: model.LeagueQuizPlease,
		}
		game.My = true

		user := model.User{
			Name:  "user name",
			Email: model.NewMaybeString("user email"),
			Phone: model.NewMaybeString("user phone"),
		}

		quizPleaseCroupierMock.EXPECT().RegisterForLottery(context.Background(), game, user).Return(0, errors.New("some error"))

		got, err := croupier.RegisterForLottery(ctx, game, user)
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
	})

	// TODO fix squiz lottery registration
	/*
		t.Run("error while registration, squiz", func(t *testing.T) {
			squizCroupierMock := mocks.NewLotteryRegistrator(t)
			croupier := New(Config{
				SquizCroupier: squizCroupierMock,
			})

			globalConfig := config.GlobalConfig{}
			globalConfig.LotteryStartsBefore = 3600
			config.UpdateGlobalConfig(globalConfig)

			time_utils.TimeNow = func() time.Time {
				return time_utils.ConvertTime("2022-01-10 15:31")
			}

			ctx := context.Background()

			game := model.Game{
				Date:     model.DateTime(time_utils.ConvertTime("2022-01-09 16:30")),
				LeagueID: pkgmodel.LeagueSquiz,
			}
			game.My = true

			user := model.User{
				Email: "user email",
				Name:  "user name",
				Phone: "user phone",
			}

			squizCroupierMock.EXPECT().RegisterForLottery(context.Background(), game, user).Return(0, errors.New("some error"))

			got, err := croupier.RegisterForLottery(ctx, game, user)
			assert.Equal(t, int32(0), got)
			assert.Error(t, err)
		})
	*/

	t.Run("ok, quiz please", func(t *testing.T) {
		quizPleaseCroupierMock := mocks.NewLotteryRegistrator(t)
		croupier := New(Config{
			QuizPleaseCroupier: quizPleaseCroupierMock,
		})

		globalConfig := config.GlobalConfig{}
		globalConfig.LotteryStartsBefore = 3600
		config.UpdateGlobalConfig(globalConfig)

		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-10 15:31")
		}

		ctx := context.Background()

		game := model.Game{
			Date:     model.DateTime(time_utils.ConvertTime("2022-01-09 16:30")),
			LeagueID: model.LeagueQuizPlease,
		}
		game.My = true

		user := model.User{
			Name:  "user name",
			Email: model.NewMaybeString("user email"),
			Phone: model.NewMaybeString("user phone"),
		}

		quizPleaseCroupierMock.EXPECT().RegisterForLottery(context.Background(), game, user).Return(100, nil)

		got, err := croupier.RegisterForLottery(ctx, game, user)
		assert.Equal(t, int32(100), got)
		assert.NoError(t, err)
	})

	// TODO fix squiz lottery registration
	/*
		t.Run("ok squiz", func(t *testing.T) {
			squizCroupierMock := mocks.NewLotteryRegistrator(t)
			croupier := New(Config{
				SquizCroupier: squizCroupierMock,
			})

			globalConfig := config.GlobalConfig{}
			globalConfig.LotteryStartsBefore = 3600
			config.UpdateGlobalConfig(globalConfig)

			time_utils.TimeNow = func() time.Time {
				return time_utils.ConvertTime("2022-01-10 15:31")
			}

			ctx := context.Background()

			game := model.Game{
				Date:     model.DateTime(time_utils.ConvertTime("2022-01-09 16:30")),
				LeagueID: pkgmodel.LeagueSquiz,
			}
			game.My = true

			user := model.User{
				Email: "user email",
				Name:  "user name",
				Phone: "user phone",
			}

			squizCroupierMock.EXPECT().RegisterForLottery(context.Background(), game, user).Return(0, nil)

			got, err := croupier.RegisterForLottery(ctx, game, user)
			assert.Equal(t, int32(0), got)
			assert.NoError(t, err)
		})
	*/
}
