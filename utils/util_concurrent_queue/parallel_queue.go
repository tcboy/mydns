package util_concurrent_queue

import (
	"sync/atomic"
	"time"

	"math/rand"

	internalqueue "code.aliyun.com/batcn/common/utils/util_concurrent_queue/internal_queue"
)

type ParallelQueue struct {
	startIndex       int32
	RingBufferList   []*internalqueue.RingBuffer
	ParallelDegree   uint64
	ParallelCapacity uint64 // 并行队列的总容量
}

func NewParallelQueueDefalut(paralleNo uint64) *ParallelQueue {
	return NewParallelQueue(paralleNo, uint64(1024))
}

func NewParallelQueue(parallelDegree uint64, RingBufferCapacity uint64) *ParallelQueue {
	ringBufferList := make([]*internalqueue.RingBuffer, parallelDegree)

	realParalleNo := internalqueue.RoundUp(parallelDegree)

	for i := 0; i < int(realParalleNo); i++ {
		ringBuffer := internalqueue.NewRingBuffer(RingBufferCapacity)
		ringBufferList[i] = ringBuffer
	}

	realRingBufferCapacity := internalqueue.RoundUp(RingBufferCapacity)
	parallelCapacity := realRingBufferCapacity * realParalleNo

	parallelQueue := &ParallelQueue{
		0,
		ringBufferList,
		realParalleNo,
		parallelCapacity,
	}

	return parallelQueue
}

func (queue *ParallelQueue) Offer(value int64) bool {
	// if value is power of 2 , then bitwise modulus: x % n == x & (n - 1).
	index := uint64(value) & (queue.ParallelDegree - 1)
	ok, _ := queue.RingBufferList[index].Offer(value)
	if !ok {
		// logs.Info("index %d is full",index)
	}
	return ok
}

func (queue *ParallelQueue) Offers(values []int64) {
	startIndex := atomic.LoadInt32(&queue.startIndex)
	size := len(values)
	var newStartIndex int

	for i := 0; i < size; i++ {
		newStartIndex = (int(startIndex) + i) % int(queue.ParallelDegree)
		queue.RingBufferList[newStartIndex].Offer(values[i])
	}

	atomic.StoreInt32(&queue.startIndex, int32(newStartIndex))
}

func (queue *ParallelQueue) Poll(modKey int64, timeout time.Duration) (interface{}, error) {
	index := rand.Intn(int(queue.ParallelDegree))
	current := queue.RingBufferList[index]
	var (
		value interface{}
		err   error
	)
	for i := 0; i < 10; i++ {
		value, err = current.Poll(timeout)
		if value == nil {
			if current.Len() == 0 {
				degree := int(queue.ParallelDegree)
				for i := 0; i < degree; i++ {
					next := queue.RingBufferList[(index+i)/degree]
					value, err := next.Poll(timeout)
					if value != nil {
						return value, err
					}
				}
			} else {
				value, err = current.Poll(timeout)
			}

		}
	}
	return value, err
}
