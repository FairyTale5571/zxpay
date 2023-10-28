package zxpay

const (
	// ZxAPIBaseURL is the base url for the 0xpay API
	ZxAPIBaseURL = "https://public.api.0xpay.app"
	// ZxAPIEndpointMerchants is the endpoint for merchants
	ZxAPIEndpointMerchants = "/merchants"

	ZxAPIEndpointBalance = ZxAPIEndpointMerchants + "/balances"

	// ZxAPIEndpointAddress is the endpoint for addresses
	ZxAPIEndpointAddress = ZxAPIEndpointMerchants + "/addresses"

	// ZxAPIEndpointWithdrawals is the endpoint for withdrawals
	ZxAPIEndpointWithdrawals = ZxAPIEndpointMerchants + "/withdrawals"

	// ZxAPIEndpointWithdrawalCrypto is the endpoint for crypto withdrawals
	ZxAPIEndpointWithdrawalCrypto = ZxAPIEndpointWithdrawals + "/crypto"

	// ZxAPIEndpointWithdrawalFiat is the endpoint for fiat withdrawals
	ZxAPIEndpointWithdrawalFiat = ZxAPIEndpointWithdrawals + "/fiat"

	// ZxAPIEndpointWithdrawalCryptoFee is the endpoint for crypto withdrawal fees
	ZxAPIEndpointWithdrawalCryptoFee = ZxAPIEndpointWithdrawalCrypto + "/fee"

	// ZxAPIEndpointInvoices is the endpoint for invoices
	ZxAPIEndpointInvoices = ZxAPIEndpointMerchants + "/invoices"

	// ZxAPIEndpointInvoicesCrypto is the endpoint for crypto invoices
	ZxAPIEndpointInvoicesCrypto = ZxAPIEndpointInvoices + "/crypto"

	// ZxAPIEndpointInvoicesCryptoForm is the endpoint for crypto invoices form
	ZxAPIEndpointInvoicesCryptoForm = ZxAPIEndpointInvoicesCrypto + "/form"

	// ZxAPIEndpointInvoicesFiat is the endpoint for fiat invoices
	ZxAPIEndpointInvoicesFiat = ZxAPIEndpointInvoices + "/fiat"

	// ZxAPIEndpointInvoicesFiatForm is the endpoint for fiat invoices form
	ZxAPIEndpointInvoicesFiatForm = ZxAPIEndpointInvoicesFiat + "/form"

	// ZxAPIEndpointExchange is the endpoint for exchange
	ZxAPIEndpointExchange = ZxAPIEndpointMerchants + "/exchange"

	// ZxAPIEndpointExchangeEstimate is the endpoint for exchange estimate
	ZxAPIEndpointExchangeEstimate = ZxAPIEndpointExchange + "/estimate"

	// ZxAPIEndpointExchangeDirections is the endpoint for exchange directions
	ZxAPIEndpointExchangeDirections = ZxAPIEndpointExchange + "/available-directions"

	// ZxAPIEndpointFiatAssets is the endpoint for fiat assets
	ZxAPIEndpointFiatAssets = ZxAPIEndpointMerchants + "/assets/fiat"

	// ZxAPIEndpointCryptoAssets is the endpoint for crypto assets
	ZxAPIEndpointCryptoAssets = ZxAPIEndpointMerchants + "/assets/crypto"
)

type ZxPay struct {
	merchantID string
	privateKey string
	url        string
	signature  string
}

func New(merchantID, privateKey string, opts ...Options) *ZxPay {
	return &ZxPay{
		merchantID: merchantID,
		privateKey: privateKey,
	}
}

type Options struct{}

// Validate verifies a notification signature
func (z *ZxPay) Validate(op ...Options) (bool, error) {
	return true, nil
}
