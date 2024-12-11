package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/conflux-fans/go-scan-sdk/constants"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Client struct {
	baseUrl string
	inner   *resty.Client
}

func NewClient(baseUrl string) *Client {
	return &Client{
		baseUrl: baseUrl,
		inner:   resty.New().SetRetryCount(3),
	}
}

func (c *Client) url(_path string) string {
	return fmt.Sprintf("%s%s", c.baseUrl, _path)
}

func (c *Client) GetAccountTransactions(account, from, to string, start, end time.Time, skip, limit int, asc bool) (*List[*Transaction], error) {

	sort := "ASC"
	if !asc {
		sort = "DESC"
	}

	queryParams := map[string]string{
		"account":      account,
		"minTimestamp": strconv.FormatInt(start.Unix(), 10),
		"maxTimestamp": strconv.FormatInt(end.Unix(), 10),
		"skip":         strconv.Itoa(skip),
		"limit":        strconv.Itoa(limit),
		"sort":         sort,
	}

	if from != "" {
		queryParams["from"] = from
	}

	if to != "" {
		queryParams["to"] = to
	}

	logrus.WithField("queryParams", queryParams).Info("[Scan Client] get account transactions")

	resp, err := c.inner.R().
		SetQueryParams(queryParams).
		SetHeader("Accept", "application/json").
		Get(c.url("/account/transactions"))
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(resp.Body()))

	var data Response[*List[*Transaction]]
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if data.Code != constants.RESPONSOE_CODE_OK {
		return nil, errors.New(fmt.Sprintf("code: %d, message: %s", data.Code, data.Message))
	}

	return data.Data, nil
}

// NOTE: this api is not offically released
func (c *Client) GetPosAccountOverview(posAddress common.Hash) (*PosAccountOverview, error) {
	queryParams := map[string]string{
		"address": posAddress.Hex(),
	}

	logrus.WithField("queryParams", queryParams).Info("[Scan Client] get pos account overview")

	resp, err := c.inner.R().
		SetQueryParams(queryParams).
		SetHeader("Accept", "application/json").
		Get(c.url("/stat/pos-account-overview"))
	if err != nil {
		return nil, err
	}

	var data Response[*PosAccountOverview]
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if data.Code != constants.RESPONSOE_CODE_OK {
		return nil, errors.New(fmt.Sprintf("code: %d, message: %s", data.Code, data.Message))
	}

	return data.Data, nil
}
