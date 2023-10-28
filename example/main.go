package main

import (
	"fmt"
	"net/http"

	"github.com/fairytale5571/zxpay"
	"github.com/gin-gonic/gin"
)

const (
	merchantID = "75e5fc6f-7bf1-4e8f-a205-1f017dda4b68"
	privateKey = "cbbe94143499a0344b0910a039d05f59"
	url        = "6900-95-158-53-98.eu.ngrok.io"
)

func main() {
	zx := zxpay.New(merchantID, privateKey)

	// Get supported fiat
	fiat, err := zx.GetBalance("UAH", "USDT", "BTC")
	if err != nil {
		panic(err)
	}
	fmt.Println(fiat)
	// Get supported crypto
	crypto, err := zx.GetListCryptoAssets()
	if err != nil {
		panic(err)
	}
	fmt.Println(crypto)

	router := gin.New()
	router.Any("/", handler)
	router.Run(":3000")

}

func handler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
