package model

// CertificateType ...
type CertificateType uint8

const (
	// CertificateTypeInvalid ...
	CertificateTypeInvalid CertificateType = iota
	// CertificateTypeFreePass ...
	CertificateTypeFreePass
	// CertificateTypeBarBillPayment ...
	CertificateTypeBarBillPayment

	// NumberOfCertificateTypes ...
	NumberOfCertificateTypes
)
