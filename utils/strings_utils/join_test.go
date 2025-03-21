package strings_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string
}

func (r *User) String() string {
	return r.Name
}

func TestUnit_StringsUtils_JoinStringers_Ok(t *testing.T) {
	t.Parallel()

	users := []*User{{Name: "John"}, {Name: "Jane"}, {Name: "Artem"}}

	assert.Equal(t, "John, Jane, Artem", JoinStringers(users, ", "))
}
