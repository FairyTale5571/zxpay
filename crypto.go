package zxpay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// GetListCryptoAssets Get List of Crypto Assets
// https://docs.0xpay.app/public-api/endpoints/basic-crypto-operations#list-all-supported-crypto-assets
func (z *ZxPay) GetListCryptoAssets() (AssetsInformation, error) {
	var res AssetsInformation
	u, _ := url.Parse(fmt.Sprintf("%s%s", ZxAPIBaseURL, "/merchants/assets/crypto"))
	ca, err := z.do(http.MethodGet, u)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(ca.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// CreatePermanentAddress Create Permanent Deposit Address
// https://docs.0xpay.app/public-api/endpoints/basic-crypto-operations#create-permanent-deposit-address
func (z *ZxPay) CreatePermanentAddress() {

}

// Send Cryptocurrency
// https://docs.0xpay.app/public-api/endpoints/basic-crypto-operations#send-cryptocurrency
func (z *ZxPay) Send() {

}

type Fee struct {
	// Value decimal fee value
	Value string `json:"value"`
}

// GetSendingFee Get Sending Fee
// https://docs.0xpay.app/public-api/endpoints/basic-crypto-operations#get-sending-fee
func (z *ZxPay) GetSendingFee() {

}

// ListSupportedCryptoAssets List all supported crypto assets
// https://docs.0xpay.app/public-api/endpoints/basic-crypto-operations#list-all-supported-crypto-assets
func (z *ZxPay) ListSupportedCryptoAssets() {

}

type Invoice struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	ToPendingImmediate bool   `json:"toPendingImmediate"`
	ClientDuration     int64  `json:"clientDuration"`
	Duration           int64  `json:"duration"`
	Meta               string `json:"meta"`
	Description        string `json:"description"`
}

func (z *ZxPay) CreatePaymentInvoice(name, mail, amount string) {

}

func (z *ZxPay) CreatePaymentInvoiceComplex() {

}

func (z *ZxPay) CretePaymentInvoiceClient() {
	z.CretePaymentInvoiceClient()
}
