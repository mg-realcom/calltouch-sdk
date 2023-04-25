package calltouch_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const DateFormat = "02/01/2006"

type Client struct {
	AccessToken string
}

func NewClient(accessToken string) *Client {
	return &Client{
		AccessToken: accessToken,
	}
}

//func (c *Client) OrdersDiary(siteID int, dateFrom time.Time, dateTo time.Time) (err error) {
//	if dateFrom.After(dateTo) {
//		return errors.New("dateFrom must be before dateTo")
//	}
//
//	u, err := c.urlBuilder("orders-diary/orders", siteID, dateFrom, dateTo)
//	if err != nil {
//		return nil, err
//	}
//
//
//	req, err := http.Get(u.String())
//	if err != nil {
//		return err
//	}
//	defer req.Body.Close()
//
//	if req.StatusCode != 200 {
//		return errors.New(fmt.Sprintf("status code: %v", req.StatusCode))
//	}
//
//	responseBody, err := io.ReadAll(req.Body)
//	if err != nil {
//		return err
//	}
//
//	var data string
//	err = json.Unmarshal(responseBody, &data)
//	if err != nil {
//		return err
//	}
//
//	return err
//}

func (c *Client) urlBuilder(method string, siteID int, dateFrom time.Time, dateTo time.Time) (u url.URL, err error) {
	if dateFrom.After(dateTo) {
		return u, errors.New("dateFrom must be before dateTo")
	}

	dateFromString := dateFrom.Format(DateFormat)
	dateToString := dateTo.Format(DateFormat)

	u = url.URL{
		Scheme: "https",
		Host:   "api.calltouch.ru",
		Path:   fmt.Sprintf("calls-service/RestAPI/%v/%s", siteID, method),
	}

	params := url.Values{}
	params.Add("clientApiId", c.AccessToken)
	params.Add("dateFrom", dateFromString)
	params.Add("dateTo", dateToString)
	params.Add("page", "1")
	params.Add("limit", "1000")
	u.RawQuery = params.Encode()

	return u, nil
}

func (c *Client) CallsDiary(siteID int, dateFrom time.Time, dateTo time.Time) (calls []Call, err error) {
	u, err := c.urlBuilder("calls-diary/calls", siteID, dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	req, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code: %v", req.StatusCode))
	}

	responseBody, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data CallReport
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return nil, err
	}

	return data.Records, nil
}

type CallReport struct {
	Page         int    `json:"page"`
	PageTotal    int    `json:"pageTotal"`
	PageSize     int    `json:"pageSize"`
	RecordsTotal int    `json:"recordsTotal"`
	Records      []Call `json:"records"`
}
