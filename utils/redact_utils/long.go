package redact_utils

import (
	"github.com/teadove/teasutils/utils/must_utils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	maxLen = 12
)

func redactRecursivly(iterableValue gjson.Result, prefix string, result []byte) []byte {
	iterableValue.ForEach(func(key, value gjson.Result) bool {
		if value.Type == gjson.String && len(value.Str) > maxLen {
			newValue := RedactWithPrefixSized(value.Str, maxLen)

			// Warning!
			// Possible panic :)
			result = must_utils.Must(sjson.SetBytes(result, prefix+key.Str, newValue))

			return true
		}

		if value.Type == gjson.JSON {
			result = redactRecursivly(value, key.Str+".", result)
		}

		return true
	})

	return result
}

func RedactLongStrings(s []byte) []byte {
	return redactRecursivly(gjson.ParseBytes(s), "", s)
}
