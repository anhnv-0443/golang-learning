package constants

import (
	"time"
)

const (
	// ConnectTimeout 10s
	ConnectTimeout = time.Second * 10

	// ConnectWaitDuration for database sleep reconnect 5s
	ConnectWaitDuration = time.Second * 5

	// ConnectReadTimeout 30s
	ConnectReadTimeout = time.Second * 30
)

const (
	// TokenLifetime 7days
	TokenLifetime = time.Hour * 7 * 24
)

// Input for ParseInt
const (
	Int64Num   = 64
	DecimalNum = 10
)
