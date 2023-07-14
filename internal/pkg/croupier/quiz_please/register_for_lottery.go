package quiz_please

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// {"success":true,"message":"168"}
type response struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

type multipartObject struct {
	name  string
	value io.Reader
}

// RegisterForLottery ...
func (c *Croupier) RegisterForLottery(ctx context.Context, game model.Game, user model.User) (int32, error) {
	multipartObjects := []multipartObject{
		{
			name:  "game_id",
			value: strings.NewReader(strconv.FormatInt(int64(game.ExternalID), 10)),
		},
		{
			name:  "LotteryPlayer[team_name]",
			value: strings.NewReader("Жизнь и Грабля"),
		},
		{
			name:  "LotteryPlayer[name]",
			value: strings.NewReader(user.Name),
		},
		{
			name:  "LotteryPlayer[email]",
			value: strings.NewReader(user.Email.Value),
		},
		{
			name:  "LotteryPlayer[phone]",
			value: strings.NewReader(user.Phone.Value),
		},
	}

	var requestBody bytes.Buffer
	var err error
	w := multipart.NewWriter(&requestBody)
	err = w.SetBoundary("123")
	if err != nil {
		return 0, fmt.Errorf("set boundary error: %w", err)
	}

	for _, multipartObject := range multipartObjects {
		var fw io.Writer
		if x, ok := multipartObject.value.(io.Closer); ok {
			defer x.Close()
		}

		if fw, err = w.CreateFormField(multipartObject.name); err != nil {
			return 0, fmt.Errorf("create form field error: %w", err)
		}

		if _, err = io.Copy(fw, multipartObject.value); err != nil {
			return 0, fmt.Errorf("copy multipart object error: %w", err)
		}
	}

	w.Close()

	req, err := http.NewRequest("POST", c.lotteryLink, &requestBody)
	if err != nil {
		return 0, fmt.Errorf("creating request error: %w", err)
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("http do request error: %w", err)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("read body error: %w", err)
	}
	defer resp.Body.Close()

	var r response
	err = json.Unmarshal(responseBody, &r)
	if err != nil {
		return 0, fmt.Errorf("response body unmarshal error: %w", err)
	}

	if r.Success {
		number, err := strconv.ParseInt(r.Message, 10, 32)
		if err != nil {
			return 0, fmt.Errorf("parse number error: %w", err)
		}

		return int32(number), nil
	}

	return 0, errors.New(r.Message)
}
