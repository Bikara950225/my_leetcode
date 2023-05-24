package heap

import "testing"

func TestMinHeap_Add(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		heap := MinHeap{}

		heap.Add(10)
		if m, _ := heap.Top(); m != 10 {
			t.Fatalf("expected 10, but got %d", m)
		}
		heap.Add(2)
		if m, _ := heap.Top(); m != 2 {
			t.Fatalf("expected 2, but got %d", m)
		}

		heap.Add(1)
		if m, _ := heap.Top(); m != 1 {
			t.Fatalf("expected 1, but got %d", m)
		}

		if m, _ := heap.Pop(); m != 1 {
			t.Fatalf("expected 1, but got %d", m)
		}
		t.Log(heap)
	})
}
