package main

import (
	"context"
	"errors"
)

func GetContextValue[T any](c context.Context, key any) (T, error) {
	var t T
	val := c.Value(key)
	if val == nil {
		return t, errors.New("key doesn't exist")
	}
	t, ok := val.(T)
	if !ok {
		return t, errors.New("unexpected value type")

	}
	return t, nil
}
