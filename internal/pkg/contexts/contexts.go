package contexts

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
)

// contextIDKey is the key used to retrieve the id from the context values.
type contextIDKey struct{}

// AddIDToContext adds an id to the context that can be used for logging.
func AddIDToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextIDKey{}, uuid.NewV4().String())
}

// GetIDFromContext returns the id of the context.
func GetIDFromContext(ctx context.Context) string {
	return cast.ToString(ctx.Value(contextIDKey{}))
}
