package util

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// GenerateUid generates a UUIDv7 and retries up to 10 times in case of error.
func GenerateUid() string {
	var uid uuid.UUID
	var err error

	for i := 0; i < 10; i++ {
		if uid, err = uuid.NewV7(); err == nil {
			return uid.String()
		}
	}

	return fmt.Sprintf("%d", time.Now().UnixNano())
}
