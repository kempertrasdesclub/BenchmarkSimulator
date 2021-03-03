package user

import (
	"bytes"
	"encoding/base32"
	"github.com/pborman/uuid"
)

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

// NewId is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// with the padding stripped off.
func (e *User) NewId() (key string, err error) {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	_, err = encoder.Write(uuid.NewRandom())
	if err != nil {
		return
	}
	err = encoder.Close()
	if err != nil {
		return
	}
	b.Truncate(26) // removes the '==' padding

	key = b.String()

	return
}
