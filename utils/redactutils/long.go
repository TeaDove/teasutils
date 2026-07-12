package redactutils

import (
	"github.com/teadove/teasutils/utils/mustutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	maxLen = 150
)

func redactRecursivly(iterableValue gjson.Result, prefix string, result []byte) []byte {
	iterableValue.ForEach(func(key, value gjson.Result) bool {
		if value.Type == gjson.String && len(value.Str) > maxLen {
			newValue := RedactWithPrefixSized(value.Str, maxLen)

			// Warning!
			// Possible panic :)
			result = mustutils.Must(sjson.SetBytes(result, prefix+key.Str, newValue))

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
