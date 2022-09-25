package dom

import (
	"crypto/rand"
	"strings"

	"github.com/oklog/ulid"
)

// Auth should contain all authentication and authorization info
// needed to execute any operation on behalf of some user.
type Auth struct {
	UserName UserName
	Admin    bool
}

func NewID() string {
	return strings.ToLower(ulid.MustNew(ulid.Now(), rand.Reader).String())
}
