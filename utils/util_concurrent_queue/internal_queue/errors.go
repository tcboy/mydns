package internal_queue

import "errors"

/**
 * Error definitions
 */
var (
	ErrorCapacity = errors.New("ERROR_CAPACITY: attempt to Create EfficientRingbuffer with invalid Capacity")
	ErrorFull     = errors.New("ERROR_FULL: attempt to Put while EfficientRingbuffer is Full")
	ErrorEmpty    = errors.New("ERROR_EMPTY: attempt to Get while EfficientRingbuffer is Empty")

	// ErrDisposed is returned when an operation is performed on a disposed
	// enqueue.
	ErrDisposed = errors.New(`enqueue: disposed`)

	// ErrTimeout is returned when an applicable enqueue operation times out.
	ErrTimeout = errors.New(`enqueue: poll timed out`)

	// ErrEmptyQueue is returned when an non-applicable enqueue operation was called
	// due to the enqueue's empty item state
	ErrEmptyQueue = errors.New(`enqueue: empty enqueue`)
)
