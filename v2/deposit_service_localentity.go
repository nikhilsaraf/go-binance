package binance

import (
	"context"
	"encoding/json"
)

// ListDepositsService fetches deposit history for local entity entries.
//
// See https://developers.binance.com/docs/wallet/travel-rule/deposit-history
type ListDepositsLocalEntityService struct {
	c                    *Client
	trId                 *string
	txId                 *string
	tranId               *string
	network              *string
	coin                 *string
	travelRuleStatus     *int  // 0:Completed,1:Pending,2:Failed
	pendingQuestionnaire *bool // true: Only return records that pending deposit questionnaire. false/not provided: return all records.
	startTime            *int64
	endTime              *int64
	offset               *int
	limit                *int
}

// TrID sets the trId parameter.
func (s *ListDepositsLocalEntityService) TrID(trId string) *ListDepositsLocalEntityService {
	s.trId = &trId
	return s
}

// TxID sets the txId parameter.
func (s *ListDepositsLocalEntityService) TxID(txId string) *ListDepositsLocalEntityService {
	s.txId = &txId
	return s
}

// TranID sets the tranId parameter.
func (s *ListDepositsLocalEntityService) TranID(tranId string) *ListDepositsLocalEntityService {
	s.tranId = &tranId
	return s
}

// Network sets the network parameter.
func (s *ListDepositsLocalEntityService) Network(network string) *ListDepositsLocalEntityService {
	s.network = &network
	return s
}

// Coin sets the coin parameter.
func (s *ListDepositsLocalEntityService) Coin(coin string) *ListDepositsLocalEntityService {
	s.coin = &coin
	return s
}

// TravelRuleStatus sets the travelRuleStatus parameter.
func (s *ListDepositsLocalEntityService) TravelRuleStatus(travelRuleStatus int) *ListDepositsLocalEntityService {
	s.travelRuleStatus = &travelRuleStatus
	return s
}

// PendingQuestionnaire sets the pendingQuestionnaire parameter.
func (s *ListDepositsLocalEntityService) PendingQuestionnaire(pendingQuestionnaire bool) *ListDepositsLocalEntityService {
	s.pendingQuestionnaire = &pendingQuestionnaire
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListDepositsLocalEntityService) StartTime(startTime int64) *ListDepositsLocalEntityService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListDepositsLocalEntityService) EndTime(endTime int64) *ListDepositsLocalEntityService {
	s.endTime = &endTime
	return s
}

// Offset set offset
func (s *ListDepositsLocalEntityService) Offset(offset int) *ListDepositsLocalEntityService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *ListDepositsLocalEntityService) Limit(limit int) *ListDepositsLocalEntityService {
	s.limit = &limit
	return s
}

// Do sends the request.
func (s *ListDepositsLocalEntityService) Do(ctx context.Context) (res []*DepositLocalEntity, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/localentity/deposit/history",
		secType:  secTypeSigned,
	}
	if s.trId != nil {
		r.setParam("trId", *s.trId)
	}
	if s.txId != nil {
		r.setParam("txId", *s.txId)
	}
	if s.tranId != nil {
		r.setParam("tranId", *s.tranId)
	}
	if s.network != nil {
		r.setParam("network", *s.network)
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.travelRuleStatus != nil {
		r.setParam("travelRuleStatus", *s.travelRuleStatus)
	}
	if s.pendingQuestionnaire != nil {
		r.setParam("pendingQuestionnaire", *s.pendingQuestionnaire)
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
	res = make([]*DepositLocalEntity, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// DepositLocalEntity represents a single DepositLocalEntity entry.
type DepositLocalEntity struct {
	TrID                 int64   `json:"trId"`
	TranID               string  `json:"tranId"`
	Amount               string  `json:"amount"`
	Coin                 string  `json:"coin"`
	Network              string  `json:"network"`
	Status               int     `json:"status"`
	TravelRuleStatus     int     `json:"travelRuleStatus"`
	DepositStatus        int     `json:"depositStatus"`
	Address              string  `json:"address"`
	AddressTag           string  `json:"addressTag"`
	TxID                 string  `json:"txId"`
	InsertTime           int64   `json:"insertTime"`
	TransferType         int64   `json:"transferType"`
	ConfirmTimes         string  `json:"confirmTimes"`
	RequireQuestionnaire bool    `json:"requireQuestionnaire"`
	Questionnaire        *string `json:"questionnaire"`
}

// SubmitDepositQuestionnaireService submits the deposit questionnaire
//
// See https://developers.binance.com/docs/wallet/travel-rule/deposit-provide-info
type SubmitDepositQuestionnaireService struct {
	c                           *Client
	tranId                      int64
	questionnaireJsonUrlEncoded string
}

// TranID sets the tranId parameter.
func (s *SubmitDepositQuestionnaireService) TranID(tranId int64) *SubmitDepositQuestionnaireService {
	s.tranId = tranId
	return s
}

// QuestionnaireJsonUrlEncoded sets the questionnaireJsonUrlEncoded parameter.
func (s *SubmitDepositQuestionnaireService) QuestionnaireJsonUrlEncoded(questionnaireJsonUrlEncoded string) *SubmitDepositQuestionnaireService {
	s.questionnaireJsonUrlEncoded = questionnaireJsonUrlEncoded
	return s
}

// Do sends the request.
func (s *SubmitDepositQuestionnaireService) Do(ctx context.Context) (*SubmitDepositQuestionnaireResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/localentity/deposit/provide-info",
		secType:  secTypeSigned,
	}

	r.setParam("tranId", s.tranId)
	r.setParam("questionnaire", s.questionnaireJsonUrlEncoded)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &SubmitDepositQuestionnaireResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SubmitDepositQuestionnaireResponse represents a response from SubmitDepositQuestionnaireService.
type SubmitDepositQuestionnaireResponse struct {
	TrID     int    `json:"trId"`
	Accepted bool   `json:"accepted"`
	Info     string `json:"info"`
}
