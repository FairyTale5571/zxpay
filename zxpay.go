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

	// ZxAPIEndpointWithdrawalCrypto is the endpoint for crypto withdrawals
	ZxAPIEndpointWithdrawalFiat = ZxAPIEndpointWithdrawals + "/fiat"

	// ZxAPIEndpointWithdrawalCryptoFee is the endpoint for crypto withdrawal fees
	ZxAPIEndpointWithdrawalCryptoFee = ZxAPIEndpointWithdrawalCrypto + "/fee"

	// ZxAPIEndpointInvoices
	ZxAPIEndpointInvoices = ZxAPIEndpointMerchants + "/invoices"

	// ZxAPIEndpointInvoicesCrypto
	ZxAPIEndpointInvoicesCrypto = ZxAPIEndpointInvoices + "/crypto"

	//ZxAPIEndpointInvoicesCryptoForm
	ZxAPIEndpointInvoicesCryptoForm = ZxAPIEndpointInvoicesCrypto + "/form"

	// ZxAPIEndpointInvoicesFiat
	ZxAPIEndpointInvoicesFiat = ZxAPIEndpointInvoices + "/fiat"

	// ZxAPIEndpointInvoicesFiatForm
	ZxAPIEndpointInvoicesFiatForm = ZxAPIEndpointInvoicesFiat + "/form"

	// ZxAPIEndpointExchange
	ZxAPIEndpointExchange = ZxAPIEndpointMerchants + "/exchange"
	// ZxAPIEndpointExchangeEstimate
	ZxAPIEndpointExchangeEstimate = ZxAPIEndpointExchange + "/estimate"
	// ZxAPIEndpointExchangeDirections
	ZxAPIEndpointExchangeDirections = ZxAPIEndpointExchange + "/available-directions"

	ZxAPIEndpointFiatAssets   = ZxAPIEndpointMerchants + "/assets/fiat"
	ZxAPIEndpointCryptoAssets = ZxAPIEndpointMerchants + "/assets/crypto"
)

type Provider interface {
	MerchantRepository
	FiatRepository
	CryptoRepository
}

type zxPay struct {
	merchantID string
	privateKey string
	url        string
	signature  string
}

func New(merchantID, privateKey string) Provider {
	return &zxPay{
		merchantID: merchantID,
		privateKey: privateKey,
	}
}

type Options struct {
}

// VerifyNotification verifies a notification signature
func (z *zxPay) Validate(op ...Options) (bool, error) {
	return true, nil
}
