package domain

import (
	"time"

	"github.com/satori/go.uuid"
)

// Tweet interfase
type Tweet interface {
	PrintableTweet() string
	GetUser() string
	GetText() string
	GetId() uuid.UUID
	GetDate() *time.Time
}
