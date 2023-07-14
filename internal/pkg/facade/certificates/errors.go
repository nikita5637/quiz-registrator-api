package certificates

import "errors"

var (
	// ErrCertificateNotFound ...
	ErrCertificateNotFound = errors.New("certificate not found")
	// ErrWonOnGameNotFound ...
	ErrWonOnGameNotFound = errors.New("won on game not found")
	// ErrSpentOnGameNotFound ...
	ErrSpentOnGameNotFound = errors.New("spent on game not found")
)
