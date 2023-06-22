package vo

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	amount   int
	form     uuid.UUID
	to       uuid.UUID
	createAt time.Time
}
