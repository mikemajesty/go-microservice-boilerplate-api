package utils

import "context"

func ContextWithValues(ctx context.Context, key any, value string) context.Context {
	return context.WithValue(ctx, key, value)
}
