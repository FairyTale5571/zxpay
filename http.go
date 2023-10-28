package zxpay

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func (z *zxPay) do(method string, path *url.URL, body ...interface{}) (*http.Response, error) {
	var req *http.Request
	var client = &http.Client{Timeout: time.Second * 10}
	var err error
	var b []byte

	if body != nil {
		b, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, path.String(), bytes.NewReader(b))
	} else {
		req, err = http.NewRequest(method, path.String(), http.NoBody)
	}
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().Unix() - 5
	signature := z.generateSignature(method, path.RequestURI(), string(b), timestamp)
	req.Header.Set("merchant-id", z.merchantID)
	req.Header.Set("signature", signature)
	req.Header.Set("timestamp", fmt.Sprintf("%d", timestamp))

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		var errorResponse APIError

		defer response.Body.Close()
		reader, _ := io.ReadAll(response.Body)
		fmt.Printf("%s\n", string(reader))

		if err := json.Unmarshal(reader, &errorResponse); err != nil {
			return nil, err
		}
		return nil, err
	}
	return response, nil
}

func (z *zxPay) generateSignature(method, url, body string, timestamp int64) string {
	s := fmt.Sprintf("%s%s%s%d", method, url, body, timestamp)
	h := hmac.New(crypto.SHA256.New, []byte(z.privateKey))
	_, err := h.Write([]byte(s))
	if err != nil {
		return ""
	}
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
