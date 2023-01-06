package quiz_please

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"

	"github.com/stretchr/testify/assert"
)

func TestQuizPleaseCroupier_RegisterForLottery(t *testing.T) {
	t.Run("error 1", func(t *testing.T) {
		croupier := QuizPleaseCroupier{}
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := ioutil.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.Equal(t,
				[]byte{0x2d, 0x2d, 0x31, 0x32, 0x33, 0xd, 0xa, 0x43, 0x6f, 0x6e,
					0x74, 0x65, 0x6e, 0x74, 0x2d, 0x44, 0x69, 0x73, 0x70, 0x6f,
					0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x66, 0x6f,
					0x72, 0x6d, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x3b, 0x20, 0x6e,
					0x61, 0x6d, 0x65, 0x3d, 0x22, 0x67, 0x61, 0x6d, 0x65, 0x5f,
					0x69, 0x64, 0x22, 0xd, 0xa, 0xd, 0xa, 0x37, 0x37, 0x37, 0xd,
					0xa, 0x2d, 0x2d, 0x31, 0x32, 0x33, 0xd, 0xa, 0x43, 0x6f, 0x6e,
					0x74, 0x65, 0x6e, 0x74, 0x2d, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73,
					0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x66, 0x6f, 0x72, 0x6d,
					0x2d, 0x64, 0x61, 0x74, 0x61, 0x3b, 0x20, 0x6e, 0x61, 0x6d, 0x65,
					0x3d, 0x22, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x50, 0x6c,
					0x61, 0x79, 0x65, 0x72, 0x5b, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e,
					0x61, 0x6d, 0x65, 0x5d, 0x22, 0xd, 0xa, 0xd, 0xa, 0xd0, 0x96, 0xd0,
					0xb8, 0xd0, 0xb7, 0xd0, 0xbd, 0xd1, 0x8c, 0x20, 0xd0, 0xb8, 0x20,
					0xd0, 0x93, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xb1, 0xd0, 0xbb, 0xd1,
					0x8f, 0xd, 0xa, 0x2d, 0x2d, 0x31, 0x32, 0x33, 0xd, 0xa, 0x43, 0x6f,
					0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x44, 0x69, 0x73, 0x70, 0x6f,
					0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x66, 0x6f, 0x72,
					0x6d, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x3b, 0x20, 0x6e, 0x61, 0x6d,
					0x65, 0x3d, 0x22, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x50,
					0x6c, 0x61, 0x79, 0x65, 0x72, 0x5b, 0x6e, 0x61, 0x6d, 0x65, 0x5d,
					0x22, 0xd, 0xa, 0xd, 0xa, 0xd0, 0x98, 0xd0, 0xbc, 0xd1, 0x8f, 0xd,
					0xa, 0x2d, 0x2d, 0x31, 0x32, 0x33, 0xd, 0xa, 0x43, 0x6f, 0x6e, 0x74,
					0x65, 0x6e, 0x74, 0x2d, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69, 0x74,
					0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x64, 0x61,
					0x74, 0x61, 0x3b, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x4c,
					0x6f, 0x74, 0x74, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65,
					0x72, 0x5b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5d, 0x22, 0xd, 0xa,
					0xd, 0xa, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0xd, 0xa, 0x2d, 0x2d, 0x31,
					0x32, 0x33, 0xd, 0xa, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d,
					0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
					0x20, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x3b, 0x20,
					0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x4c, 0x6f, 0x74, 0x74, 0x65, 0x72,
					0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5b, 0x70, 0x68, 0x6f, 0x6e,
					0x65, 0x5d, 0x22, 0xd, 0xa, 0xd, 0xa, 0xd0, 0x9d, 0xd0, 0xbe, 0xd0,
					0xbc, 0xd0, 0xb5, 0xd1, 0x80, 0x20, 0xd1, 0x82, 0xd0, 0xb5, 0xd0, 0xbb,
					0xd0, 0xb5, 0xd1, 0x84, 0xd0, 0xbe, 0xd0, 0xbd, 0xd0, 0xb0, 0xd, 0xa, 0x2d,
					0x2d, 0x31, 0x32, 0x33, 0x2d, 0x2d, 0xd, 0xa}, b)
			fmt.Fprintf(w, `{"success":false,"message":""}`)
		}))
		defer svr.Close()

		croupier.lotteryLink = svr.URL
		got, err := croupier.RegisterForLottery(context.Background(), model.Game{
			ExternalID: 777,
		}, model.User{
			Name:  "Имя",
			Email: "Email",
			Phone: "Номер телефона",
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.Equal(t, "", err.Error())
	})

	t.Run("error 2", func(t *testing.T) {
		croupier := QuizPleaseCroupier{}
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, `{"success":false,"message":"Игра не найдена"}`)
		}))
		defer svr.Close()

		croupier.lotteryLink = svr.URL
		got, err := croupier.RegisterForLottery(context.Background(), model.Game{
			ExternalID: 777,
		}, model.User{})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.Equal(t, "Игра не найдена", err.Error())
	})

	t.Run("error 3", func(t *testing.T) {
		croupier := QuizPleaseCroupier{}
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, `{"success":false,"message":"Кажется, игрок с таким емейлом уже зарегистрирован в лототроне сегодня.<br><br> <p>Сорри, мы не можем принять вашу запись (но вы можете предложить вашему сокоманднику поучаствовать).</p>"}`)
		}))
		defer svr.Close()

		croupier.lotteryLink = svr.URL
		got, err := croupier.RegisterForLottery(context.Background(), model.Game{
			ExternalID: 777,
		}, model.User{})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.Equal(t, "Кажется, игрок с таким емейлом уже зарегистрирован в лототроне сегодня.<br><br> <p>Сорри, мы не можем принять вашу запись (но вы можете предложить вашему сокоманднику поучаствовать).</p>", err.Error())
	})

	t.Run("error 4", func(t *testing.T) {
		croupier := QuizPleaseCroupier{}
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, `{"success":false,"message":"Лотерея не найдена"}`)
		}))
		defer svr.Close()

		croupier.lotteryLink = svr.URL
		got, err := croupier.RegisterForLottery(context.Background(), model.Game{
			ExternalID: 777,
		}, model.User{})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.Equal(t, "Лотерея не найдена", err.Error())
	})

	t.Run("ok", func(t *testing.T) {
		croupier := QuizPleaseCroupier{}
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, `{"success":true,"message":"168"}`)
		}))
		defer svr.Close()

		croupier.lotteryLink = svr.URL
		got, err := croupier.RegisterForLottery(context.Background(), model.Game{
			ExternalID: 777,
		}, model.User{})
		assert.Equal(t, int32(168), got)
		assert.NoError(t, err)
	})

}
