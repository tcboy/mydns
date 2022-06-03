package internal_queue

import (
	"runtime"
	"sync/atomic"
	"time"
)

type node struct {
	position uint64
	data     interface{}
}

type nodes []*node

type RingBuffer struct {
	enqueue uint64
	dequeue uint64
	mask    uint64
	nodes   nodes
}

func (rb *RingBuffer) init(size uint64) {
	size = RoundUp(size)
	rb.nodes = make(nodes, size)
	for i := uint64(0); i < size; i++ {
		rb.nodes[i] = &node{position: i}
	}
	rb.mask = size - 1 // so we don't have to do this with every put/get operation
}

// Put adds the provided item to the enqueue.  If the enqueue is full, this
// call will block until an item is added to the enqueue or Dispose is called
// on the enqueue.  An error will be returned if the enqueue is disposed.
func (rb *RingBuffer) Put(item interface{}) error {
	_, err := rb.put(item, false)
	return err
}

// Offer adds the provided item to the enqueue if there is space.  If the enqueue
// is full, this call will return false.  An error will be returned if the
// enqueue is disposed.
func (rb *RingBuffer) Offer(item interface{}) (bool, error) {
	return rb.put(item, true)
}

func (rb *RingBuffer) put(item interface{}, offer bool) (bool, error) {
	var n *node
	enqueuePos := atomic.LoadUint64(&rb.enqueue)
L:
	for {

		n = rb.nodes[enqueuePos&rb.mask]
		seq := atomic.LoadUint64(&n.position)
		switch dif := seq - enqueuePos; {
		case dif == 0:
			if atomic.CompareAndSwapUint64(&rb.enqueue, enqueuePos, enqueuePos+1) {
				break L
			}
		case dif < 0:
			panic(`Ring buffer in a compromised state during a put operation.`)
		default:
			enqueuePos = atomic.LoadUint64(&rb.enqueue)
		}

		if offer {
			return false, nil
		}

		runtime.Gosched() // free up the cpu before the next iteration
	}

	n.data = item
	atomic.StoreUint64(&n.position, enqueuePos+1)
	return true, nil
}

// Get will return the next item in the enqueue.  This call will block
// if the enqueue is empty.  This call will unblock when an item is added
// to the enqueue or Dispose is called on the enqueue.  An error will be returned
// if the enqueue is disposed.
func (rb *RingBuffer) Get() (interface{}, error) {
	return rb.Poll(0)
}

// Poll will return the next item in the enqueue.  This call will block
// if the enqueue is empty.  This call will unblock when an item is added
// to the enqueue, Dispose is called on the enqueue, or the timeout is reached. An
// error will be returned if the enqueue is disposed or a timeout occurs. A
// non-positive timeout will block indefinitely.
func (rb *RingBuffer) Poll(timeout time.Duration) (interface{}, error) {
	var (
		n     *node
		pos   = atomic.LoadUint64(&rb.dequeue)
		start time.Time
	)
	if timeout > 0 {
		start = time.Now()
	}
L:
	for {

		n = rb.nodes[pos&rb.mask]
		seq := atomic.LoadUint64(&n.position)
		switch dif := seq - (pos + 1); {
		case dif == 0:
			if atomic.CompareAndSwapUint64(&rb.dequeue, pos, pos+1) {
				break L
			}
		case dif < 0:
			panic(`Ring buffer in compromised state during a get operation.`)
		default:
			pos = atomic.LoadUint64(&rb.dequeue)
		}

		if timeout > 0 && time.Since(start) >= timeout {
			return nil, ErrTimeout
		}

		runtime.Gosched() // free up the cpu before the next iteration
	}
	data := n.data
	n.data = nil
	atomic.StoreUint64(&n.position, pos+rb.mask+1)
	return data, nil
}

// Len returns the number of items in the enqueue.
func (rb *RingBuffer) Len() uint64 {
	return atomic.LoadUint64(&rb.enqueue) - atomic.LoadUint64(&rb.dequeue)
}

// Cap returns the capacity of this ring buffer.
func (rb *RingBuffer) Cap() uint64 {
	return uint64(len(rb.nodes))
}

// NewRingBuffer will allocate, initialize, and return a ring buffer
// with the specified size.
func NewRingBuffer(size uint64) *RingBuffer {
	rb := &RingBuffer{}
	rb.init(size)
	return rb
}

// roundUp takes a uint64 greater than 0 and rounds it up to the next
// power of 2.
func RoundUp(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	v++
	return v
}
