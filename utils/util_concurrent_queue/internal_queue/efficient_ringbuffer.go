// esQueue
package internal_queue

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

// lock free enqueue
type EfficientRingbuffer struct {
	capacity         uint32
	capacityRoundOf2 uint32
	writeIndex       uint32
	readIndex        uint32
	valueWrappers    []ValueWrapper
}

type ValueWrapper struct {
	wOffset uint32 //总是递增
	rOffset uint32 //总是递增
	value   interface{}
}

func NewEfficientRingbuffer(capaciity uint32) *EfficientRingbuffer {
	queue := new(EfficientRingbuffer)
	queue.capacity = RoundForTwo(capaciity)
	queue.capacityRoundOf2 = queue.capacity - 1
	queue.writeIndex = 0
	queue.readIndex = 0
	queue.valueWrappers = make([]ValueWrapper, queue.capacity)
	for i := range queue.valueWrappers {
		cache := &queue.valueWrappers[i]
		cache.rOffset = uint32(i)
		cache.wOffset = uint32(i)
	}
	cache := &queue.valueWrappers[0]
	cache.rOffset = queue.capacity
	cache.wOffset = queue.capacity
	return queue
}

func (queue *EfficientRingbuffer) String() string {
	readIndex := atomic.LoadUint32(&queue.readIndex)
	writeIndex := atomic.LoadUint32(&queue.writeIndex)
	return fmt.Sprintf("EfficientRingbuffer{capacity: %v, capacityRoundOf2: %v, writeIndex: %v, readIndex: %v}",
		queue.capacity, queue.capacityRoundOf2, writeIndex, readIndex)
}

func (queue *EfficientRingbuffer) Capacity() uint32 {
	return queue.capacity
}

func (queue *EfficientRingbuffer) ElementsNo() uint32 {
	var wi, ri uint32
	var elementsNo uint32
	ri = atomic.LoadUint32(&queue.readIndex)
	wi = atomic.LoadUint32(&queue.writeIndex)

	if wi >= ri {
		elementsNo = wi - ri
	} else {
		elementsNo = queue.capacityRoundOf2 + (wi - ri)
	}

	return elementsNo
}

// non-blocking 达到capMod-1,立即返回,
func (queue *EfficientRingbuffer) Put(val interface{}) (ok bool, available uint32) {
	var writeIndex, newWriteIndex, readIndex, elementsNo uint32
	var valueWrapper *ValueWrapper
	capMod := queue.capacityRoundOf2

	for ok := true; ok; { //cas do-while

		readIndex = atomic.LoadUint32(&queue.readIndex)
		writeIndex = atomic.LoadUint32(&queue.writeIndex)

		if writeIndex >= readIndex {
			elementsNo = writeIndex - readIndex
		} else {
			elementsNo = capMod + (writeIndex - readIndex)
		}

		if elementsNo >= capMod-1 {
			return false, elementsNo
		}

		newWriteIndex = writeIndex + 1
		if atomic.CompareAndSwapUint32(&queue.writeIndex, writeIndex, newWriteIndex) {
			ok = false
		} else {
			runtime.Gosched()
		}

	}
	//increment_address_one = (address + 1) % Length
	// decrement_address_one = (address + Length -1) % Length
	// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
	valueWrapper = &queue.valueWrappers[newWriteIndex&capMod]

	for {
		rOffset := atomic.LoadUint32(&valueWrapper.rOffset)
		wOffset := atomic.LoadUint32(&valueWrapper.wOffset)
		if newWriteIndex == wOffset && rOffset == wOffset {
			valueWrapper.value = val
			atomic.AddUint32(&valueWrapper.wOffset, queue.capacity)
			return true, elementsNo + 1
		} else {
			runtime.Gosched()
		}
	}
}

// get enqueue functions
func (queue *EfficientRingbuffer) Get() (val interface{}, ok bool, quantity uint32) {
	var writeIndex, readIndex, readIndexNew, elementNo uint32
	var cache *ValueWrapper
	capMod := queue.capacityRoundOf2

	writeIndex = atomic.LoadUint32(&queue.writeIndex)
	readIndex = atomic.LoadUint32(&queue.readIndex)

	if writeIndex >= readIndex {
		elementNo = writeIndex - readIndex
	} else {
		elementNo = capMod + (writeIndex - readIndex)
	}

	if elementNo < 1 {
		runtime.Gosched()
		return nil, false, elementNo
	}

	readIndexNew = readIndex + 1
	if !atomic.CompareAndSwapUint32(&queue.readIndex, readIndex, readIndexNew) {
		runtime.Gosched()
		return nil, false, elementNo
	}

	cache = &queue.valueWrappers[readIndexNew&capMod]

	for {
		getNo := atomic.LoadUint32(&cache.rOffset)
		putNo := atomic.LoadUint32(&cache.wOffset)
		if readIndexNew == getNo && getNo == putNo-queue.capacity {
			val = cache.value
			cache.value = nil
			atomic.AddUint32(&cache.rOffset, queue.capacity)
			return val, true, elementNo - 1
		} else {
			runtime.Gosched()
		}
	}
}

// puts enqueue functions
func (queue *EfficientRingbuffer) Puts(values []interface{}) (puts, quantity uint32) {
	var putPos, putPosNew, getPos, posCnt, putCnt uint32
	capMod := queue.capacityRoundOf2

	getPos = atomic.LoadUint32(&queue.readIndex)
	putPos = atomic.LoadUint32(&queue.writeIndex)

	if putPos >= getPos {
		posCnt = putPos - getPos
	} else {
		posCnt = capMod + (putPos - getPos)
	}

	if posCnt >= capMod-1 {
		runtime.Gosched()
		return 0, posCnt
	}

	if capPuts, size := queue.capacity-posCnt, uint32(len(values)); capPuts >= size {
		putCnt = size
	} else {
		putCnt = capPuts
	}
	putPosNew = putPos + putCnt

	if !atomic.CompareAndSwapUint32(&queue.writeIndex, putPos, putPosNew) {
		runtime.Gosched()
		return 0, posCnt
	}

	for posNew, v := putPos+1, uint32(0); v < putCnt; posNew, v = posNew+1, v+1 {
		var cache *ValueWrapper = &queue.valueWrappers[posNew&capMod]
		for {
			getNo := atomic.LoadUint32(&cache.rOffset)
			putNo := atomic.LoadUint32(&cache.wOffset)
			if posNew == putNo && getNo == putNo {
				cache.value = values[v]
				atomic.AddUint32(&cache.wOffset, queue.capacity)
				break
			} else {
				runtime.Gosched()
			}
		}
	}
	return putCnt, posCnt + putCnt
}

// round 到最近的2的倍数
func RoundForTwo(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

func Delay(z int) {
	for x := z; x > 0; x-- {
	}
}
