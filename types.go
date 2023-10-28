package zxpay

import "sort"

// AssetInformation all crypto assets available on 0xPay
type AssetInformation struct {
	Name        string            `json:"name"`
	Ticker      string            `json:"ticker"`
	Blockchains []string          `json:"blockchains,omitempty"`
	Price       TickerInformation `json:"price"`
}

// AssetsInformation array of AssetInformation
type AssetsInformation []AssetInformation

func (a AssetsInformation) GetTickers() []string {
	var res []string
	for _, v := range a {
		res = append(res, v.Ticker)
	}
	return res
}

func (a AssetsInformation) in(ticker string) bool {
	sorted := sort.Search(len(a), func(i int) bool { return ticker <= a[i].Ticker })
	if sorted < len(a) && a[sorted].Ticker == ticker {
		return true
	}
	return false
}

// TickerInformation ticker information
type TickerInformation struct {
	Value  string `json:"value"`
	Ticker string `json:"ticker"`
}

// AccountBalances account balances
type AccountBalances []TickerInformation

type Company struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Amount struct {
	Value      string `json:"value"`
	Ticker     string `json:"ticker"`
	Blockchain string `json:"blockchain"`
}

type PaymentStatus string

const (
	StatusCreated    PaymentStatus = "CREATED"
	StatusRegistered PaymentStatus = "REGISTERED"
	StatusPending    PaymentStatus = "PENDING"
	StatusPaid       PaymentStatus = "PAID"
	StatusExpired    PaymentStatus = "EXPIRED"
)

func (p PaymentStatus) IsValid() bool {
	switch p {
	case StatusCreated, StatusRegistered, StatusPending, StatusPaid, StatusExpired:
		return true
	}
	return false
}

type BaseInvoice struct {
	// ID
	ID string `json:"id"`
	// Company
	Company *Company `json:"company"`
	// Name
	Name string `json:"name"`
	// Status
	Status PaymentStatus `json:"status"`
	// CreatedAt
	CreatedAt int64 `json:"createdAt"`
	// ExpiredAt
	ExpiredAt int64 `json:"expiredAt"`
	// Amount
	Amount *Amount `json:"amount"`
	// Fee
	Fee string `json:"fee,omitempty"`
}

// OverPaidInvoice overpaid invoice
type OverPaidInvoice struct {
	BaseInvoice

	// MerchantID
	MerchantID string `json:"merchantId"`
	// PaidAmount
	PaidAmount string `json:"paidAmount"`
	// Meta
	Meta string `json:"meta"`
}

// AmountInvoice with base amount
type AmountInvoice struct {
	BaseInvoice

	// Address
	Address string `json:"address"`
	// ClientExpiredAt
	ClientExpiredAt int64 `json:"clientExpiredAt"`
	// Price
	Price string `json:"price"`
	// PaidPrice
	PaidPrice string `json:"paidPrice"`
	// Email
	Email string `json:"email"`
	// BaseAmount
	BaseAmount TickerInformation `json:"baseAmount"`
}

// Target target
type Target struct {
}
