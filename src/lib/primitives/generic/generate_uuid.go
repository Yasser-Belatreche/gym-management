package generic

import (
	"github.com/gofrs/uuid"
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func GenerateULID() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	return ulid.MustNew(ulid.Now(), entropy).String()
}

func GenerateUUID() string {
	id, _ := uuid.NewV7()

	return id.String()
}
