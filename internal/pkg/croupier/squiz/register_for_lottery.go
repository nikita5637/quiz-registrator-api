package squiz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

type param struct {
	Key   string
	Value string
}

type params []param

func (p *params) Add(key, value string) {
	*p = append(*p, param{
		Key:   key,
		Value: value,
	})
}

func (p params) Encode() string {
	buf := strings.Builder{}
	for _, param := range p {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		if param.Key == "phone" || param.Key == "tildaspec-phone-part[]" {
			buf.WriteString(url.QueryEscape(param.Key))
			buf.WriteString("=")
			buf.WriteString(param.Value)
		} else {
			buf.WriteString(url.QueryEscape(param.Key))
			buf.WriteString("=")
			buf.WriteString(url.QueryEscape(param.Value))
		}

	}

	return buf.String()
}

type response struct {
	Message string   `json:"message,omitempty"`
	Results []string `json:"results,omitempty"`
}

// RegisterForLottery ...
func (c *Croupier) RegisterForLottery(ctx context.Context, _ model.Game, user model.User) (int32, error) {
	lotteryInfoPageReq, err := http.NewRequest(http.MethodGet, c.lotteryInfoPageLink, nil)
	if err != nil {
		return 0, fmt.Errorf("creating lottery info page request error: %w", err)
	}

	lotteryInfoPageResp, err := c.client.Do(lotteryInfoPageReq)
	if err != nil {
		return 0, fmt.Errorf("getting lottery info page error: %w", err)
	}

	doc, err := goquery.NewDocumentFromResponse(lotteryInfoPageResp)
	if err != nil {
		return 0, fmt.Errorf("can't read document: %w", err)
	}

	data := params{}
	formServices, err := c.getFormServices(ctx, doc)
	if err != nil {
		return 0, fmt.Errorf("getting form services error: %w", err)
	}
	data.Add("formservices[]", formServices)

	city, err := c.getCity(ctx, doc)
	if err != nil {
		return 0, fmt.Errorf("getting city error: %w", err)
	}

	data.Add("city", city)

	data.Add("name", user.Name)
	data.Add("team", "Жизнь и Грабля")
	data.Add("tildaspec-phone-part[]", getPhonePart(user.Phone.Value()))
	data.Add("phone", getPhone(user.Phone.Value()))
	data.Add("email", user.Email.Value())
	data.Add("policy", "yes")
	data.Add("form-spec-comments", "")
	gaCid, err := c.getGaCid(ctx, doc)
	if err != nil {
		return 0, fmt.Errorf("getting ga-cid error: %w", err)
	}
	data.Add("ga-cid", gaCid)

	ymCid, err := c.getYmCid(ctx, doc)
	if err != nil {
		return 0, fmt.Errorf("getting ym-cid error: %w", err)
	}
	data.Add("ym-cid", ymCid)

	data.Add("tildaspec-cookie", `_fbp=fb.1.1674837243537.1860351820; previousUrl=spb.squiz.ru%2Fgame; tildasid=1674837245102.474240; tildauid=1674837245102.202611; tmr_detect=0%7C1674837245889; _ga=GA1.2.2095662112.1674837243; _gid=GA1.2.858509434.1674837243; _ym_d=1674837243; _ym_isad=2; _ym_uid=1668581346714918248; _ym_visorc=w; tmr_lvid=8837cd2f89f7e2366a01121f7183882b; tmr_lvidTS=1668581346211; metrika_enabled=1`)
	data.Add("tildaspec-referer", "https://spb.squiz.ru/game")

	formID, err := c.getFormID(ctx, doc)
	if err != nil {
		return 0, fmt.Errorf("getting form ID error: %w", err)
	}
	data.Add("tildaspec-formid", formID)

	formsKey, err := c.getFormsKey(ctx, doc)
	if err != nil {
		return 0, fmt.Errorf("getting forms key error: %w", err)
	}
	data.Add("tildaspec-formskey", formsKey)

	data.Add("tildaspec-version-lib", "02.001")

	pageID, err := c.getPageID(ctx, doc)
	if err != nil {
		return 0, fmt.Errorf("getting page ID error: %w", err)
	}
	data.Add("tildaspec-pageid", pageID)

	projectID, err := c.getProjectID(ctx, doc)
	if err != nil {
		return 0, fmt.Errorf("getting project ID error: %w", err)
	}
	data.Add("tildaspec-projectid", projectID)

	data.Add("tildaspec-lang", "EN")
	data.Add("tildaspec-fp", "6354646d6863386c656e2d4742704d6163496e74656c764170706c6520436f6d70757465722c20496e632e614d6f7a696c6c616e4e65747363617065706c707232773134343068323731")

	// send registration POST request
	registrationReq, err := http.NewRequest(http.MethodPost, c.lotteryRegistrationLink, strings.NewReader(data.Encode()))
	if err != nil {
		return 0, fmt.Errorf("creating registration request error: %w", err)
	}

	registrationReq.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	registrationReq.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")

	registrationResp, err := c.client.Do(registrationReq)
	if err != nil {
		return 0, fmt.Errorf("registration request error: %w", err)
	}

	registrationRespHTML, err := ioutil.ReadAll(registrationResp.Body)
	if err != nil {
		return 0, fmt.Errorf("read registration response body error: %w", err)
	}

	response := &response{}
	err = json.Unmarshal(registrationRespHTML, response)
	if err != nil {
		return 0, fmt.Errorf("unmarshaling response error: %w", err)
	}

	if response.Message != "OK" {
		return 0, fmt.Errorf("lottery registration error: %v", response)
	}

	return 0, nil
}

func (c *Croupier) getFormServices(ctx context.Context, doc *goquery.Document) (string, error) {
	formServices := c.getValueByName(ctx, doc, "formservices[]")

	if formServices == "" {
		return "", errors.New("form services not found")
	}

	return formServices, nil
}

func (c *Croupier) getCity(ctx context.Context, doc *goquery.Document) (string, error) {
	city := c.getValueByName(ctx, doc, "city")

	if city == "" {
		return "", errors.New("city not found")
	}

	return city, nil
}

func (c *Croupier) getGaCid(_ context.Context, _ *goquery.Document) (string, error) {
	return "2095662112.1674837243", nil
}

func (c *Croupier) getYmCid(_ context.Context, _ *goquery.Document) (string, error) {
	return "1668581346714918248", nil
}

func (c *Croupier) getValueByName(_ context.Context, doc *goquery.Document, name string) string {
	var ret string
	doc.Find("form").Each(func(_ int, s *goquery.Selection) {
		doc.Find("input").Each(func(_ int, s *goquery.Selection) {
			n, ok := s.Attr("name")
			if ok && n == name {
				value, ok := s.Attr("value")
				if ok {
					ret = value
					return
				}
			}
		})
	})

	return ret
}

func (c *Croupier) getFormID(_ context.Context, doc *goquery.Document) (string, error) {
	var formID string
	doc.Find("form").Each(func(i int, s *goquery.Selection) {
		if v, ok := s.Attr("id"); ok {
			formID = v
			return
		}
	})

	if formID == "" {
		return "", errors.New("form ID not found")
	}

	return formID, nil
}

func (c *Croupier) getFormsKey(_ context.Context, doc *goquery.Document) (string, error) {
	var formsKey string
	doc.Find("#allrecords").Each(func(i int, s *goquery.Selection) {
		if v, ok := s.Attr("data-tilda-formskey"); ok {
			formsKey = v
			return
		}
	})

	if formsKey == "" {
		return "", errors.New("forms key not found")
	}

	return formsKey, nil
}

func (c *Croupier) getPageID(_ context.Context, doc *goquery.Document) (string, error) {
	var pageID string
	doc.Find("#allrecords").Each(func(i int, s *goquery.Selection) {
		if v, ok := s.Attr("data-tilda-page-id"); ok {
			pageID = v
			return
		}
	})

	if pageID == "" {
		return "", errors.New("forms key not found")
	}

	return pageID, nil
}

func (c *Croupier) getProjectID(_ context.Context, doc *goquery.Document) (string, error) {
	var projectID string
	doc.Find("#allrecords").Each(func(i int, s *goquery.Selection) {
		if v, ok := s.Attr("data-tilda-project-id"); ok {
			projectID = v
			return
		}
	})

	if projectID == "" {
		return "", errors.New("forms key not found")
	}

	return projectID, nil
}

func getPhone(phone string) string {
	return fmt.Sprintf("%s+(%s)+%s-%s-%s", url.QueryEscape(phone[0:2]), phone[2:5], phone[5:8], phone[8:10], phone[10:12])
}

func getPhonePart(phone string) string {
	return fmt.Sprintf("(%s)+%s-%s-%s", phone[2:5], phone[5:8], phone[8:10], phone[10:12])
}
