package json_utils

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog"
)

func MarshalOrWarn(ctx context.Context, v any) []byte {
	marshalled, err := json.Marshal(v)
	if err == nil {
		return marshalled
	}

	zerolog.Ctx(ctx).
		Warn().
		Stack().Err(err).
		Interface("v", v).
		Msg("failed to marshal json")

	return []byte("{}")
}
