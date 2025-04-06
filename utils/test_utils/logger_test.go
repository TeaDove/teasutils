package test_utils

import "testing"

func TestLogAny(t *testing.T) {
	t.Parallel()

	Pprint("1", 1, "РАЗ-ДВА-ТРИ", map[string]string{"a": "b", "c": "d"}, nil, func() {})
}
