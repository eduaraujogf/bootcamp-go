package main

import (
	"context"
	"fmt"
)

type contextKey int

const ContextKeyID contextKey = iota

func greetingsWrapper(ctx context.Context) {
	ctx = context.WithValue(ctx, ContextKeyID, "World")
	greetings(ctx)
}

func greetings(ctx context.Context) {
	fmt.Println(ctx.Value(ContextKeyID))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextKeyID, "Hello")

	greetingsWrapper(ctx)
}
