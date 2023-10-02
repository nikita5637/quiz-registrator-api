package authentication

import (
	"context"
)

type serviceAuthType struct{}

var (
	serviceAuthKey = serviceAuthType{}
)

// NewContextWithServiceAuth ...
func NewContextWithServiceAuth(ctx context.Context) context.Context {
	return context.WithValue(ctx, serviceAuthKey, true)
}

// IsServiceAuth ...
func IsServiceAuth(ctx context.Context) bool {
	val, ok := ctx.Value(serviceAuthKey).(bool)
	if !ok {
		return false
	}

	return val
}
