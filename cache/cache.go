package cache

import (
	"errors"
	"time"
)

const (
	NoneDuration = time.Duration(-1)
)

var ErrKeyNotFound = errors.New("cache key not found")

type Factory interface {
	UserCaches() UserCache
	MessageCaches() MessageCache
}