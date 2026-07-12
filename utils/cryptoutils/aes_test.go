package cryptoutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/teadove/teasutils/utils/testutils"
)

func TestAES(t *testing.T) {
	t.Parallel()

	key := []byte(TextWithLen(32))
	msg := "Hello World!"

	encrypted, err := AESEncrypt([]byte(msg), key)
	require.NoError(t, err)

	testutils.Pprint(encrypted)

	decrypted, err := AESDecrypt(encrypted, key)
	require.NoError(t, err)

	assert.Equal(t, msg, string(decrypted))
}
