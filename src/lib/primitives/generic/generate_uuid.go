package generic

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func GenerateUUID() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	return ulid.MustNew(ulid.Now(), entropy).String()
}
