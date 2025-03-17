package strings_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string
}

func (r *User) String() string {
	return r.Name
}

func TestUnit_StringsUtils_JoinStringers_Ok(t *testing.T) {
	users := []*User{{Name: "John"}, {Name: "Jane"}, {Name: "Artem"}}

	assert.Equal(t, "John, Jane, Artem", JoinStringers(users, ", "))
}
