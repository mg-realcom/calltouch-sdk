package calltouch

import (
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

const DateFormat = "01/02/2006"

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

func (c *Client) CallsDiary(ctx context.Context, siteID int, period Period, options *CallOptions) ([]Record, error) {
	page := 0
	calls := make([]Record, 0)

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
		isOk = data.PageTotal == data.Page

		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Println(closeErr)
		}
	}

	return calls, nil
}

func (c *Client) callURLBuilder(method string, siteID int, period Period, page int, options *CallOptions) (url.URL, error) {
	if period.DateFrom.After(period.DateTo) {
		return url.URL{}, errors.New("dateFrom must be before dateTo")
	}

	dateFromString := period.DateFrom.Format(DateFormat)
	dateToString := period.DateTo.Format(DateFormat)

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

	if options != nil {
		params.Add("uniqueOnly", strconv.FormatBool(options.UniqueOnly))
		params.Add("targetOnly", strconv.FormatBool(options.TargetOnly))
		params.Add("uniqTargetOnly", strconv.FormatBool(options.UniqTargetOnly))
		params.Add("callbackOnly", strconv.FormatBool(options.CallbackOnly))
		params.Add("withMapVisits", strconv.FormatBool(options.WithMapVisits))
		params.Add("withOrders", strconv.FormatBool(options.WithOrders))
		params.Add("withCallTags", strconv.FormatBool(options.WithCallTags))
		params.Add("withComments", strconv.FormatBool(options.WithComments))
		params.Add("withYandexDirect", strconv.FormatBool(options.WithYandexDirect))
		params.Add("withGoogleAdwords", strconv.FormatBool(options.WithGoogleAdwords))
		params.Add("withText", strconv.FormatBool(options.WithText))
		params.Add("withDcm", strconv.FormatBool(options.WithDcm))
	}

	u.RawQuery = params.Encode()

	return u, nil
}

type CallReport struct {
	Page         int      `json:"page"`
	PageTotal    int      `json:"pageTotal"`
	PageSize     int      `json:"pageSize"`
	RecordsTotal int      `json:"recordsTotal"`
	Records      []Record `json:"records"`
}

type Record map[string]string

func (r *Record) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}

	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	var conv string

	result := make(map[string]string, 0)

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
	withMapVisits     bool // Флаг истории посещений посетителя, совершившего звонок.
	withRequestTags   bool // Флаг выгрузки тегов, которые были присвоены заявкам.
	withYandexDirect  bool // Флаг выгрузки данных по рекламным кампаниям Яндекс.Директ.
	withGoogleAdwords bool // Флаг выгрузки данных по рекламным кампаниям Google AdWords.
	withDcm           bool // Флаг выгрузки данных по интеграции с DoubleClick Campaign Manager
}

func (c *Client) LeadsDiary(ctx context.Context, period Period, options *LeadOptions) ([]Record, error) {
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

	var leads []Record

	err = json.Unmarshal(responseBody, &leads)
	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (c *Client) leadURLBuilder(period Period, options *LeadOptions) (url.URL, error) {
	if period.DateFrom.After(period.DateTo) {
		return url.URL{}, errors.New("dateFrom must be before dateTo")
	}

	dateFromString := period.DateFrom.Format(DateFormat)
	dateToString := period.DateTo.Format(DateFormat)

	u := url.URL{
		Scheme: "https",
		Host:   "api.calltouch.ru",
		Path:   "calls-service/RestAPI/requests/",
	}

	params := url.Values{}
	params.Add("clientApiId", c.accessToken)
	params.Add("dateFrom", dateFromString)
	params.Add("dateTo", dateToString)

	if options != nil {
		params.Add("withMapVisits", strconv.FormatBool(options.withMapVisits))
		params.Add("withRequestTags", strconv.FormatBool(options.withRequestTags))
		params.Add("withYandexDirect", strconv.FormatBool(options.withYandexDirect))
		params.Add("withGoogleAdwords", strconv.FormatBool(options.withGoogleAdwords))
		params.Add("withDcm", strconv.FormatBool(options.withDcm))
	}

	u.RawQuery = params.Encode()

	return u, nil
}
