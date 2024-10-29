package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateWithdrawService submits a withdraw request.
//
// See https://binance-docs.github.io/apidocs/spot/en/#withdraw
type CreateWithdrawService struct {
	c                           *Client
	coin                        string
	withdrawOrderID             *string
	network                     *string
	address                     string
	addressTag                  *string
	amount                      string
	transactionFeeFlag          *bool
	questionnaireJsonUrlEncoded *string
	name                        *string
}

// Coin sets the coin parameter (MANDATORY).
func (s *CreateWithdrawService) Coin(v string) *CreateWithdrawService {
	s.coin = v
	return s
}

// WithdrawOrderID sets the withdrawOrderID parameter.
func (s *CreateWithdrawService) WithdrawOrderID(v string) *CreateWithdrawService {
	s.withdrawOrderID = &v
	return s
}

// Network sets the network parameter.
func (s *CreateWithdrawService) Network(v string) *CreateWithdrawService {
	s.network = &v
	return s
}

// Address sets the address parameter (MANDATORY).
func (s *CreateWithdrawService) Address(v string) *CreateWithdrawService {
	s.address = v
	return s
}

// AddressTag sets the addressTag parameter.
func (s *CreateWithdrawService) AddressTag(v string) *CreateWithdrawService {
	s.addressTag = &v
	return s
}

// Amount sets the amount parameter (MANDATORY).
func (s *CreateWithdrawService) Amount(v string) *CreateWithdrawService {
	s.amount = v
	return s
}

// TransactionFeeFlag sets the transactionFeeFlag parameter.
func (s *CreateWithdrawService) TransactionFeeFlag(v bool) *CreateWithdrawService {
	s.transactionFeeFlag = &v
	return s
}

// Name sets the name parameter.
func (s *CreateWithdrawService) Name(v string) *CreateWithdrawService {
	s.name = &v
	return s
}

// QuestionnaireJsonUrlEncoded sets the name parameter.
func (s *CreateWithdrawService) QuestionnaireJsonUrlEncoded(v string) *CreateWithdrawService {
	s.questionnaireJsonUrlEncoded = &v
	return s
}

// Do sends the request.
func (s *CreateWithdrawService) Do(ctx context.Context) (*CreateWithdrawResponse, error) {
	endpoint := "/sapi/v1/capital/withdraw/apply"
	if s.questionnaireJsonUrlEncoded != nil {
		endpoint = "/sapi/v1/localentity/withdraw/apply"
	}
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	r.setParam("coin", s.coin)
	r.setParam("address", s.address)
	r.setParam("amount", s.amount)
	if v := s.withdrawOrderID; v != nil {
		r.setParam("withdrawOrderId", *v)
	}
	if v := s.network; v != nil {
		r.setParam("network", *v)
	}
	if v := s.addressTag; v != nil {
		r.setParam("addressTag", *v)
	}
	if v := s.transactionFeeFlag; v != nil {
		r.setParam("transactionFeeFlag", *v)
	}
	if v := s.name; v != nil {
		r.setParam("name", *v)
	}
	if v := s.questionnaireJsonUrlEncoded; v != nil {
		r.setParam("questionnaire", *v)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	if s.questionnaireJsonUrlEncoded != nil {
		res := &CreateLocalEntityWithdrawResponse{}
		if err := json.Unmarshal(data, res); err != nil {
			return nil, err
		}
		return &CreateWithdrawResponse{
			ID: fmt.Sprintf("%d", res.ID),
		}, nil
	}

	res := &CreateWithdrawResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateWithdrawResponse represents a response from CreateWithdrawService.
type CreateWithdrawResponse struct {
	ID string `json:"id"`
}

// CreateLocalEntityWithdrawResponse represents a response from CreateWithdrawService.
type CreateLocalEntityWithdrawResponse struct {
	ID       int    `json:"trId"`
	Accepted bool   `json:"accepted"`
	Info     string `json:"info"`
}

// ListWithdrawsService fetches withdraw history.
//
// See https://binance-docs.github.io/apidocs/spot/en/#withdraw-history-supporting-network-user_data
type ListWithdrawsService struct {
	c         *Client
	coin      *string
	status    *int
	startTime *int64
	endTime   *int64
	offset    *int
	limit     *int
}

// Coin sets the coin parameter.
func (s *ListWithdrawsService) Coin(coin string) *ListWithdrawsService {
	s.coin = &coin
	return s
}

// Status sets the status parameter.
func (s *ListWithdrawsService) Status(status int) *ListWithdrawsService {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListWithdrawsService) StartTime(startTime int64) *ListWithdrawsService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListWithdrawsService) EndTime(endTime int64) *ListWithdrawsService {
	s.endTime = &endTime
	return s
}

// Offset set offset
func (s *ListWithdrawsService) Offset(offset int) *ListWithdrawsService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *ListWithdrawsService) Limit(limit int) *ListWithdrawsService {
	s.limit = &limit
	return s
}

// Do sends the request.
func (s *ListWithdrawsService) Do(ctx context.Context) (res []*Withdraw, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/withdraw/history",
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = make([]*Withdraw, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

func (s *ListWithdrawsService) DoLocalEntity(ctx context.Context) (res []*WithdrawLocalEntity, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/localentity/withdraw/history",
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = make([]*WithdrawLocalEntity, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// Withdraw represents a single withdraw entry.
type Withdraw struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ApplyTime       string `json:"applyTime"`
	Coin            string `json:"coin"`
	ID              string `json:"id"`
	WithdrawOrderID string `json:"withdrawOrderID"`
	Network         string `json:"network"`
	TransferType    int    `json:"transferType"`
	Status          int    `json:"status"`
	TransactionFee  string `json:"transactionFee"`
	TxID            string `json:"txId"`
}

type WithdrawLocalEntity struct {
	Address          string `json:"address"`
	Amount           string `json:"amount"`
	ApplyTime        int64  `json:"applyTime"`
	Coin             string `json:"coin"`
	CompleteTime     int64  `json:"completeTime"`
	ConfirmNo        int    `json:"confirmNo"`
	WithdrawalStatus int    `json:"withdrawalStatus"`
	ID               string `json:"id"`
	Info             string `json:"info"`
	Network          string `json:"network"`
	Questionnaire    struct {
		IsAddressOwner int    `json:"isAddressOwner"`
		BnfType        int    `json:"bnfType"`
		BnfName        string `json:"bnfName"`
		Country        string `json:"country"`
		SendTo         int    `json:"sendTo"`
	} `json:"questionnaire"`
	TravelID         int    `json:"trId"`
	TransferType     int    `json:"transferType"`
	TravelRuleStatus int    `json:"travelRuleStatus"`
	TransactionFee   string `json:"transactionFee"`
	TxID             string `json:"txKey"`
}
