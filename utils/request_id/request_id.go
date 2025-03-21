package request_id

import (
	"context"
	"github.com/teadove/teasutils/utils/logger_utils"
	"github.com/teadove/teasutils/utils/random_utils"
	"github.com/teadove/teasutils/utils/settings_utils"
	"strings"
)

const (
	Header    = "x-request-id"
	ctxKey    = "request_id"
	prefixLen = 8
	maxLen    = 50
)

var suffix = makeSuffix() //nolint: gochecknoglobals // Required

func makeSuffix() string {
	var builder strings.Builder

	builder.WriteByte('-')

	for _, part := range strings.Split(settings_utils.ServiceSettings.ServiceName, "-") {
		if part == "" {
			continue
		}

		builder.WriteByte(part[0])
	}

	return strings.ToUpper(builder.String())
}

func set(ctx context.Context, v string) context.Context {
	return logger_utils.WithReadableValue(ctx, ctxKey, v)
}

func MakeIfEmpty(ctx context.Context, id string) (context.Context, string) {
	if id == "" {
		return Make(ctx)
	}

	if len(id) < maxLen {
		id += suffix
	}

	return set(ctx, id), id
}

func Make(ctx context.Context) (context.Context, string) {
	id := random_utils.StringWithLen(prefixLen) + suffix
	return set(ctx, id), id
}

func GetOrMake(ctx context.Context) (context.Context, string) {
	id := logger_utils.ReadValue(ctx, ctxKey)
	if id == "" {
		return Make(ctx)
	}

	return ctx, id
}
