package limit

import (
	"sync"
	"testing"
	"time"
)

func TestTokenBucket_Check(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		tb := NewTokenBucket(50, 2*time.Second)
		wg := sync.WaitGroup{}
		wg.Add(100)

		testFunc := func(n int) {
			defer wg.Done()

			ok := tb.Check()
			t.Logf("goroutine[%d] check: %t\n", n, ok)
		}

		for i := 0; i < 100; i++ {
			go testFunc(i)
		}
		wg.Wait()

		time.Sleep(3 * time.Second)

		wg.Add(100)
		for i := 100; i < 200; i++ {
			go testFunc(i)
		}
		wg.Wait()
	})
}
