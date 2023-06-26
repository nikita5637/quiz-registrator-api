package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// CertificateType ...
type CertificateType uint8

const (
	// CertificateTypeInvalid ...
	CertificateTypeInvalid CertificateType = iota
	// CertificateTypeFreePass ...
	CertificateTypeFreePass
	// CertificateTypeBarBillPayment ...
	CertificateTypeBarBillPayment

	numberOfCertificateTypes
)

// ToSQL ...
func (t CertificateType) ToSQL() uint8 {
	return uint8(t)
}

// ValidateCertificateType ...
func ValidateCertificateType(value interface{}) error {
	v, ok := value.(CertificateType)
	if !ok {
		return errors.New("must be CertificateType")
	}

	return validation.Validate(v, validation.Max(numberOfCertificateTypes-1))
}

// Certificate ...
type Certificate struct {
	ID      int32
	Type    CertificateType
	WonOn   int32
	SpentOn MaybeInt32
	Info    MaybeString
}
