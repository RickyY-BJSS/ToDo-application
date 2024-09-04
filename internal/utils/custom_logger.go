package utils

import (
	"context"
	"log"
)

const traceIDKey ContextKey = "traceID"

func LogWithTraceID(ctx context.Context, message string) {
	traceID := ctx.Value(traceIDKey)
    log.Printf("Trace ID: %v - %s", traceID, message)
}