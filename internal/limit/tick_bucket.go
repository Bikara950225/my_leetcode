package limit

import (
	"sync"
	"time"
)

type TickBucket struct {
	duration time.Duration
	last     time.Time
	l        sync.Mutex
}

func NewTickBucket(duration time.Duration) *TokenBucket {
	return &TokenBucket{
		duration: duration,
	}
}

func (s *TickBucket) check() bool {
	s.l.Lock()
	defer s.l.Unlock()

	now := time.Now()
	if s.last.IsZero() {
		s.last = now
	}

	if now.Sub(s.last) > s.duration {
		s.last = now
		return true
	} else {
		return false
	}
}
