package zxpay

import "time"

type CreateCryptoInvoiceRequest struct {
	Email              string            `json:"email"`
	ToPendingImmediate bool              `json:"toPendingImmediate"`
	ClientDuration     time.Duration     `json:"clientDuration"`
	Duration           time.Duration     `json:"duration"`
	BaseAmount         TickerInformation `json:"baseAmount"`
	Meta               string            `json:"meta"`
	Amount             Amount            `json:"amount"`
	Description        string            `json:"description"`
	Name               string            `json:"name"`
	RedirectURL        string            `json:"redirectUrl"`
	Target             string            `json:"target"`
}
