package zxpay

import "errors"

var (
	ErrorTickerArrayEmpty   = errors.New("ticker array can not be empty")
	ErrorTickerNotSupported = errors.New("ticker: %s is unsupported")
)
