package zxpay

import (
	"net/http"
	"testing"
)

func Test_zxPay_generateSignature(t *testing.T) {
	type args struct {
		method    string
		url       string
		body      string
		timestamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all ok",
			args: args{
				method:    http.MethodGet,
				url:       "http://localhost",
				body:      "",
				timestamp: 3332211,
			},
			want: "87d373924738e85fbd13b4a9932276e894eaea4ac46092464199ca1713dea306",
		},
	}
	z := &zxPay{
		privateKey: "123",
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := z.generateSignature(tt.args.method, tt.args.url, tt.args.body, tt.args.timestamp); got != tt.want {
				t.Errorf("generateSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
