package usecase

import "errors"

var (
	// payment
	ErrBlockedPayment          = errors.New("selected payment is blocked by admin")
	ErrPaymentAmountReachedMax = errors.New("order total price reached payment method maximum amount")
	ErrPaymentNotApproved      = errors.New("payment not approved")
)
