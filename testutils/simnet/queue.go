package simnet

import (
	"sync"
)

type queue struct {
	mu        sync.Mutex
	packets   packetHeap
	newPacket chan struct{}
	closed    bool
	pushCount int
}

func newQueue() *queue { _ = "STUB: not implemented"; return nil }

func (q *queue) Enqueue(p *packetWithDeliveryTime) { _ = "STUB: not implemented"; return }

func (q *queue) Dequeue() (*packetWithDeliveryTime, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (q *queue) Close() { _ = "STUB: not implemented"; return }

type packetWithDeliveryTimeAndOrder struct {
	count int
	*packetWithDeliveryTime
}

type packetHeap []packetWithDeliveryTimeAndOrder

func (h packetHeap) Len() int { _ = "STUB: not implemented"; return 0 }

func (h packetHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (h packetHeap) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (h *packetHeap) Push(x any) { _ = "STUB: not implemented"; return }

func (h *packetHeap) Pop() any { _ = "STUB: not implemented"; return *new(any) }
