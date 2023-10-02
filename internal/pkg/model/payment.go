package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Payment ...
type Payment int32

const (
	// PaymentCash ...
	PaymentCash Payment = iota + 1
	// PaymentCertificate ...
	PaymentCertificate
	// PaymentMixed ...
	PaymentMixed

	numberOfPayments
)

// ValidatePayment ...
func ValidatePayment(value interface{}) error {
	v, ok := value.(Payment)
	if !ok {
		return errors.New("must be Payment")
	}

	return validation.Validate(v, validation.Required, validation.Min(PaymentCash), validation.Max(numberOfPayments-1))
}
