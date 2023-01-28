package squiz

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestSquizCroupier_RegisterForLottery(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()
		croupier := Croupier{}

		svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reader := strings.NewReader(html1)
			_, err := io.Copy(w, reader)
			assert.NoError(t, err)
		}))
		defer svr1.Close()

		svr2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := ioutil.ReadAll(r.Body)
			assert.NoError(t, err)

			ct := r.Header["Content-Type"]
			assert.Len(t, ct, 1)
			assert.Equal(t, "application/x-www-form-urlencoded; charset=UTF-8", ct[0])

			a := r.Header["Accept"]
			assert.Len(t, a, 1)
			assert.Equal(t, "application/json, text/javascript, */*; q=0.01", a[0])

			assert.Equal(t, `formservices%5B%5D=2f10f04603e10edcc2a2f556014f4647&city=%D0%A1%D0%B0%D0%BD%D0%BA%D1%82-%D0%9F%D0%B5%D1%82%D0%B5%D1%80%D0%B1%D1%83%D1%80%D0%B3&name=%D0%90%D0%BB%D1%91%D0%BD%D0%B0&team=%D0%96%D0%B8%D0%B7%D0%BD%D1%8C+%D0%B8+%D0%93%D1%80%D0%B0%D0%B1%D0%BB%D1%8F&tildaspec-phone-part%5B%5D=(963)+129-98-23&phone=%2B7+(963)+129-98-23&email=pilot.yes%40mail.ru&policy=yes&form-spec-comments=&ga-cid=2095662112.1674837243&ym-cid=1668581346714918248&tildaspec-cookie=_fbp%3Dfb.1.1674837243537.1860351820%3B+previousUrl%3Dspb.squiz.ru%252Fgame%3B+tildasid%3D1674837245102.474240%3B+tildauid%3D1674837245102.202611%3B+tmr_detect%3D0%257C1674837245889%3B+_ga%3DGA1.2.2095662112.1674837243%3B+_gid%3DGA1.2.858509434.1674837243%3B+_ym_d%3D1674837243%3B+_ym_isad%3D2%3B+_ym_uid%3D1668581346714918248%3B+_ym_visorc%3Dw%3B+tmr_lvid%3D8837cd2f89f7e2366a01121f7183882b%3B+tmr_lvidTS%3D1668581346211%3B+metrika_enabled%3D1&tildaspec-referer=https%3A%2F%2Fspb.squiz.ru%2Fgame&tildaspec-formid=form89161196&tildaspec-formskey=e1f3c41503ecee331fd3e20a0956db51&tildaspec-version-lib=02.001&tildaspec-pageid=4917751&tildaspec-projectid=1057010&tildaspec-lang=EN&tildaspec-fp=6354646d6863386c656e2d4742704d6163496e74656c764170706c6520436f6d70757465722c20496e632e614d6f7a696c6c616e4e65747363617065706c707232773134343068323731`,
				string(b))

			fmt.Fprintf(w, `{"message": "OK","results": ["1057010:4240991651"]}`)
		}))
		defer svr2.Close()

		croupier.lotteryInfoPageLink = svr1.URL
		croupier.lotteryRegistrationLink = svr2.URL

		got, err := croupier.RegisterForLottery(ctx, model.Game{
			ID:       1,
			LeagueID: 2,
		}, model.User{
			Name:  "Алёна",
			Phone: "+79631299823",
			Email: "pilot.yes@mail.ru",
		})
		assert.Equal(t, int32(0), got)
		assert.NoError(t, err)
	})
}

func TestSquizCroupier_getFormServices(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reader := strings.NewReader(html1)
			_, err := io.Copy(w, reader)
			assert.NoError(t, err)
		}))
		defer svr1.Close()

		croupier := New(Config{
			LotteryInfoPageLink: svr1.URL,
		})

		req, err := http.NewRequest(http.MethodGet, croupier.lotteryInfoPageLink, nil)
		assert.NoError(t, err)

		resp, err := croupier.client.Do(req)
		assert.NoError(t, err)

		doc, err := goquery.NewDocumentFromResponse(resp)
		assert.NoError(t, err)

		got, err := croupier.getFormServices(ctx, doc)
		assert.Equal(t, "2f10f04603e10edcc2a2f556014f4647", got)
		assert.NoError(t, err)
	})
}

func TestSquizCroupier_getCity(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reader := strings.NewReader(html1)
			_, err := io.Copy(w, reader)
			assert.NoError(t, err)
		}))
		defer svr1.Close()

		croupier := New(Config{
			LotteryInfoPageLink: svr1.URL,
		})

		req, err := http.NewRequest(http.MethodGet, croupier.lotteryInfoPageLink, nil)
		assert.NoError(t, err)

		resp, err := croupier.client.Do(req)
		assert.NoError(t, err)

		doc, err := goquery.NewDocumentFromResponse(resp)
		assert.NoError(t, err)

		got, err := croupier.getCity(ctx, doc)
		assert.Equal(t, "Санкт-Петербург", got)
		assert.NoError(t, err)
	})
}

func TestSquizCroupier_getGaCid(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reader := strings.NewReader(html1)
			_, err := io.Copy(w, reader)
			assert.NoError(t, err)
		}))
		defer svr1.Close()

		croupier := New(Config{
			LotteryInfoPageLink: svr1.URL,
		})

		req, err := http.NewRequest(http.MethodGet, croupier.lotteryInfoPageLink, nil)
		assert.NoError(t, err)

		resp, err := croupier.client.Do(req)
		assert.NoError(t, err)

		doc, err := goquery.NewDocumentFromResponse(resp)
		assert.NoError(t, err)

		got, err := croupier.getGaCid(ctx, doc)
		assert.Equal(t, "2095662112.1674837243", got)
		assert.NoError(t, err)
	})
}

func TestSquizCroupier_getYmCid(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reader := strings.NewReader(html1)
			_, err := io.Copy(w, reader)
			assert.NoError(t, err)
		}))
		defer svr1.Close()

		croupier := New(Config{
			LotteryInfoPageLink: svr1.URL,
		})

		req, err := http.NewRequest(http.MethodGet, croupier.lotteryInfoPageLink, nil)
		assert.NoError(t, err)

		resp, err := croupier.client.Do(req)
		assert.NoError(t, err)

		doc, err := goquery.NewDocumentFromResponse(resp)
		assert.NoError(t, err)

		got, err := croupier.getYmCid(ctx, doc)
		assert.Equal(t, "1668581346714918248", got)
		assert.NoError(t, err)
	})
}

func TestSquizCroupier_getFormID(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reader := strings.NewReader(html1)
			_, err := io.Copy(w, reader)
			assert.NoError(t, err)
		}))
		defer svr1.Close()

		croupier := New(Config{
			LotteryInfoPageLink: svr1.URL,
		})

		req, err := http.NewRequest(http.MethodGet, croupier.lotteryInfoPageLink, nil)
		assert.NoError(t, err)

		resp, err := croupier.client.Do(req)
		assert.NoError(t, err)

		doc, err := goquery.NewDocumentFromResponse(resp)
		assert.NoError(t, err)

		got, err := croupier.getFormID(ctx, doc)
		assert.Equal(t, "form89161196", got)
		assert.NoError(t, err)
	})
}

func TestSquizCroupier_getFormsKey(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reader := strings.NewReader(html1)
			_, err := io.Copy(w, reader)
			assert.NoError(t, err)
		}))
		defer svr1.Close()

		croupier := New(Config{
			LotteryInfoPageLink: svr1.URL,
		})

		req, err := http.NewRequest(http.MethodGet, croupier.lotteryInfoPageLink, nil)
		assert.NoError(t, err)

		resp, err := croupier.client.Do(req)
		assert.NoError(t, err)

		doc, err := goquery.NewDocumentFromResponse(resp)
		assert.NoError(t, err)

		got, err := croupier.getFormsKey(ctx, doc)
		assert.Equal(t, "e1f3c41503ecee331fd3e20a0956db51", got)
		assert.NoError(t, err)
	})
}

func TestSquizCroupier_getPageID(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reader := strings.NewReader(html1)
			_, err := io.Copy(w, reader)
			assert.NoError(t, err)
		}))
		defer svr1.Close()

		croupier := New(Config{
			LotteryInfoPageLink: svr1.URL,
		})

		req, err := http.NewRequest(http.MethodGet, croupier.lotteryInfoPageLink, nil)
		assert.NoError(t, err)

		resp, err := croupier.client.Do(req)
		assert.NoError(t, err)

		doc, err := goquery.NewDocumentFromResponse(resp)
		assert.NoError(t, err)

		got, err := croupier.getPageID(ctx, doc)
		assert.Equal(t, "4917751", got)
		assert.NoError(t, err)
	})
}

func TestSquizCroupier_getProjectID(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reader := strings.NewReader(html1)
			_, err := io.Copy(w, reader)
			assert.NoError(t, err)
		}))
		defer svr1.Close()

		croupier := New(Config{
			LotteryInfoPageLink: svr1.URL,
		})

		req, err := http.NewRequest(http.MethodGet, croupier.lotteryInfoPageLink, nil)
		assert.NoError(t, err)

		resp, err := croupier.client.Do(req)
		assert.NoError(t, err)

		doc, err := goquery.NewDocumentFromResponse(resp)
		assert.NoError(t, err)

		got, err := croupier.getProjectID(ctx, doc)
		assert.Equal(t, "1057010", got)
		assert.NoError(t, err)
	})
}

func Test_getPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				phone: "+79998887766",
			},
			want: "%2B7+(999)+888-77-66",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPhone(tt.args.phone); got != tt.want {
				t.Errorf("getPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPhonePart(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				phone: "+79998887766",
			},
			want: "(999)+888-77-66",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPhonePart(tt.args.phone); got != tt.want {
				t.Errorf("getPhonePart() = %v, want %v", got, tt.want)
			}
		})
	}
}
