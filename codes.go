package zxpay

type ErrorCode int

type InvoicesCode int

const (
	CodeFutureRequest    ErrorCode = -3
	CodeSignatureExpired ErrorCode = -2
	CodeInvalidSignature ErrorCode = -1

	// CodeLocalIDAlreadyPlaced - Local id already placed
	CodeLocalIDAlreadyPlaced ErrorCode = iota - 2
	// CodeNotEnoughBalance - Not enough balance
	CodeNotEnoughBalance
	// CodeFeeLessThanRequired - Fee less than required
	CodeFeeLessThanRequired
	// CodeWithdrawAmountTooLow - Withdraw amount too low
	CodeWithdrawAmountTooLow
	// CodeMerchantNotFound - Merchant not found
	CodeMerchantNotFound
	// CodeBalanceNotFound - Balance not found
	CodeBalanceNotFound
	// CodeYouDontHavePermissionForThis - You don't have permission for this
	CodeYouDontHavePermissionForThis
)

const (
	// CodeCreated - Created
	CodeCreated InvoicesCode = iota + 1
	// CodeRegistered - Registered
	CodeRegistered
	// CodePending - Pending
	CodePending
	// CodePaid - Paid
	CodePaid
	// CodeExpired - Expired
	CodeExpired
)

type Error interface {
	Error() string
}

type APIError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	StatusCode  string `json:"status_code,omitempty"`
}

func (e APIError) Error() string {
	return e.Description
}
