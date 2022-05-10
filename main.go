package main

import (
	"context"
	"fmt"
	"reflect"
)

type contextKey string

const userIDKey contextKey = "user_id"

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, userIDKey, 1)
	userID, err := GetContextValue[int](ctx, userIDKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userID)
	fmt.Println(reflect.TypeOf(userID))

	userIDStr, err := GetContextValue[string](ctx, userIDKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userIDStr)
	fmt.Println(reflect.TypeOf(userIDStr))
}
