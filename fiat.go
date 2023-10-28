package zxpay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type FiatRepository interface {
	GetListFiatAssets() (AssetsInformation, error)
	GetFiatFee(amount float64, ticker string) (*Fee, error)
}

// GetListFiatAssets return all available fiat information
func (z *zxPay) GetListFiatAssets() (AssetsInformation, error) {
	u, _ := url.Parse(fmt.Sprintf("%s%s", ZxAPIBaseURL, "/merchants/withdrawals/fiat/assets"))
	req, err := z.do(http.MethodGet, u)
	if err != nil {
		return nil, err
	}
	var res AssetsInformation
	if err := json.NewDecoder(req.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetFiatFee return fee value for fiat currency
func (z *zxPay) GetFiatFee(amount float64, ticker string) (*Fee, error) {

	listFiat, err := z.GetListFiatAssets()
	if err != nil {
		return nil, err
	}
	if !listFiat.in(ticker) {
		return nil, fmt.Errorf(ErrorTickerNotSupported.Error(), ticker)
	}
	u, _ := url.Parse(fmt.Sprintf("%s%s", ZxAPIBaseURL, "/merchants/withdrawals/fiat/fee"))
	q := u.Query()
	q.Set("amount", fmt.Sprintf("%f", amount))
	q.Set("ticker", ticker)
	u.RawQuery = q.Encode()

	fee, err := z.do(http.MethodGet, u)
	if err != nil {
		return nil, err
	}
	var res Fee
	if err := json.NewDecoder(fee.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
