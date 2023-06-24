package limit

import (
	"sync"
	"time"
)

type TokenBucket struct {
	cap      int
	count    int
	duration time.Duration
	last     time.Time // millisecond

	l sync.Mutex
}

func NewTokenBucket(cap int, duration time.Duration) *TokenBucket {
	return &TokenBucket{
		cap: cap, duration: duration,
		count: cap,
	}
}

func (s *TokenBucket) Check() bool {
	s.l.Lock()
	defer s.l.Unlock()

	now := time.Now()
	if s.last.IsZero() {
		s.last = now
	}

	// 先根据当前事件和上一次取token的时间，补充bucket
	addBuckets := now.Sub(s.last) / s.duration
	if addBuckets > 0 {
		if s.count+int(addBuckets) >= s.cap {
			s.count = s.cap
		} else {
			s.count += int(addBuckets)
		}
	}

	// 之后再判断是否可以获取桶
	if s.count > 0 {
		s.last = now
		s.count--
		return true
	} else {
		return false
	}
}
