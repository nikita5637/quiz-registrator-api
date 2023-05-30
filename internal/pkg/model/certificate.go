package model

import pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"

// Certificate ...
type Certificate struct {
	ID      int32
	Type    pkgmodel.CertificateType
	WonOn   int32
	SpentOn MaybeInt32
	Info    MaybeString
}
