package calltouch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	LeadsDateFormat = "01/02/2006"
	CallsDataFormat = "02/01/2006"
)

type Client struct {
	accessToken string
}

func NewClient(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
	}
}

type Period struct {
	DateFrom time.Time
	DateTo   time.Time
}

type CallOptions struct {
	UniqueOnly        bool // Флаг выгрузки уникальных звонков.
	TargetOnly        bool // Флаг выгрузки целевых звонков.
	UniqTargetOnly    bool // Флаг выгрузки уникально-целевых звонков.
	CallbackOnly      bool // Флаг выгрузки обратных звонков.
	WithMapVisits     bool // Флаг истории посещений посетителя, совершившего звонок.
	WithOrders        bool // Флаг связки звонка со сделкой.
	WithCallTags      bool // Флаг отображения тегов звонков.
	WithComments      bool // Флаг отображения комментариев к звонкам, оставленных в плеере журнала звонков.
	WithYandexDirect  bool // Флаг выгрузки данных по рекламным кампаниям Яндекс.Директ.
	WithGoogleAdwords bool // Флаг выгрузки данных по рекламным кампаниям Google AdWords.
	WithText          bool // Флаг выгрузки разговора звонка в текстовой форме.
	WithDcm           bool // Флаг выгрузки данных по интеграции с DoubleClick Campaign Manager
}

func (c *Client) CallsDiary(ctx context.Context, siteID int, period Period, options map[string]bool) ([]Call, error) {
	page := 0
	calls := make([]Call, 0)
	rawCalls := make([]string, 0)

	var isOk bool
	for !isOk {
		page++

		u, bErr := c.callURLBuilder("calls-diary/calls", siteID, period, page, options)
		if bErr != nil {
			return nil, bErr
		}

		req, reqErr := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
		if reqErr != nil {
			return nil, reqErr
		}

		resp, respErr := http.DefaultClient.Do(req)
		if respErr != nil {
			return nil, respErr
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status code: %v", resp.StatusCode)
		}

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var data CallReport

		jErr := json.Unmarshal(responseBody, &data)
		if err != nil {
			return nil, jErr
		}

		calls = append(calls, data.Records...)
		rawCalls = append(rawCalls, string(responseBody))
		isOk = data.PageTotal == data.Page

		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Println(closeErr)
		}
	}

	return calls, nil
}

func (c *Client) callURLBuilder(method string, siteID int, period Period, page int, options map[string]bool) (url.URL, error) {
	if period.DateFrom.After(period.DateTo) {
		return url.URL{}, errors.New("dateFrom must be before dateTo")
	}

	dateFromString := period.DateFrom.Format(CallsDataFormat)
	dateToString := period.DateTo.Format(CallsDataFormat)

	u := url.URL{
		Scheme: "https",
		Host:   "api.calltouch.ru",
		Path:   fmt.Sprintf("calls-service/RestAPI/%v/%s", siteID, method),
	}

	params := url.Values{}
	params.Add("clientApiId", c.accessToken)
	params.Add("dateFrom", dateFromString)
	params.Add("dateTo", dateToString)
	params.Add("page", strconv.Itoa(page))
	params.Add("limit", "1000")

	for k, v := range options {
		params.Add(k, strconv.FormatBool(v))
	}
	u.RawQuery = params.Encode()

	return u, nil
}

type CallReport struct {
	Page         int    `json:"page"`
	PageTotal    int    `json:"pageTotal"`
	PageSize     int    `json:"pageSize"`
	RecordsTotal int    `json:"recordsTotal"`
	Records      []Call `json:"records"`
}

type Record map[string]any

func (r *Record) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}

	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	var conv string

	result := make(map[string]any, 0)

	for k, val := range raw {
		if val == nil {
			continue
		}

		switch val := val.(type) {
		case string:
			if val == "" {
				continue
			}

			conv = strings.TrimFunc(val, func(r rune) bool { return !unicode.IsGraphic(r) })
			conv = strings.Map(func(r rune) rune {
				if unicode.IsPrint(r) {
					return r
				}
				return -1
			}, conv)
		case int:
			conv = fmt.Sprintf("%d", val)
		case float64:
			conv = fmt.Sprintf("%f", val)
		default:
			conv = fmt.Sprintf("%v", val)
		}

		result[k] = conv
	}

	*r = result

	return nil
}

type LeadOptions struct {
	WithMapVisits     bool // Флаг истории посещений посетителя, совершившего звонок.
	WithRequestTags   bool // Флаг выгрузки тегов, которые были присвоены заявкам.
	WithYandexDirect  bool // Флаг выгрузки данных по рекламным кампаниям Яндекс.Директ.
	WithGoogleAdwords bool // Флаг выгрузки данных по рекламным кампаниям Google AdWords.
	WithDcm           bool // Флаг выгрузки данных по интеграции с DoubleClick Campaign Manager
}

func (c *Client) LeadsDiary(ctx context.Context, period Period, options map[string]bool) ([]Lead, error) {
	u, err := c.leadURLBuilder(period, options)
	if err != nil {
		return nil, err
	}

	req, reqErr := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if reqErr != nil {
		return nil, reqErr
	}

	resp, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		return nil, respErr
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code - %v, reason - %v", resp.StatusCode, resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var leads []Lead

	err = json.NewDecoder(bytes.NewReader(responseBody)).Decode(&leads)
	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (c *Client) leadURLBuilder(period Period, options map[string]bool) (url.URL, error) {
	if period.DateFrom.After(period.DateTo) {
		return url.URL{}, errors.New("dateFrom must be before dateTo")
	}

	dateFromString := period.DateFrom.Format(LeadsDateFormat)
	dateToString := period.DateTo.Format(LeadsDateFormat)

	u := url.URL{
		Scheme: "https",
		Host:   "api.calltouch.ru",
		Path:   "calls-service/RestAPI/requests/",
	}

	params := url.Values{}
	params.Add("clientApiId", c.accessToken)
	params.Add("dateFrom", dateFromString)
	params.Add("dateTo", dateToString)

	for k, v := range options {
		params.Add(k, strconv.FormatBool(v))
	}
	u.RawQuery = params.Encode()
	return u, nil
}
