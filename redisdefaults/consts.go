package redisdefaults

import (
	"time"

	"gitlab.com/evatix-go/core/constants"
)

const (
	Address               = "localhost:6379"
	EmptyPassword         = constants.EmptyString
	FirstDatabase         = 0
	Localhost             = "localhost"
	Port                  = 6379
	DefaultIdleTime       = constants.N5 * time.Minute
	DefaultFrequencyCheck = constants.N5 * time.Minute
)
