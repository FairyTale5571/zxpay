package zxpay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// GetBalance return all balance from merchant
func (z *ZxPay) GetBalance(tickers ...string) (AccountBalances, error) {
	var res AccountBalances
	if len(tickers) < 1 {
		return nil, ErrorTickerArrayEmpty
	}
	query := SplitArray(tickers)
	u, e := url.Parse(fmt.Sprintf("%s%s", ZxAPIBaseURL, "/merchants/balances"))
	if e != nil {
		fmt.Errorf("%w", e)
	}
	q := u.Query()
	q.Set("tickers", query)
	u.RawQuery = q.Encode()

	req, err := z.do(http.MethodGet, u)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(req.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetFiatBalance return fiat balance from merchant
func (z *ZxPay) GetFiatBalance() (AccountBalances, error) {
	listFiat, err := z.GetListFiatAssets()
	if err != nil {
		return nil, err
	}
	tickers := listFiat.GetTickers()
	return z.GetBalance(tickers...)
}

// GetCryptoBalance return crypto balance from merchant
func (z *ZxPay) GetCryptoBalance() (AccountBalances, error) {
	listCrypto, err := z.GetListCryptoAssets()
	if err != nil {
		return nil, err
	}
	tickers := listCrypto.GetTickers()
	return z.GetBalance(tickers...)
}
