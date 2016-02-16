package ring

// Ring is a growing FIFO ring queue
type Ring struct {
	off    int
	size   int
	buffer []interface{}
}

// Add a new item onto the queue. Will grow the buffer
// if its full
func (r *Ring) Add(e interface{}) {
	if r.size < len(r.buffer) {
		end := (r.off + r.size) % len(r.buffer)
		r.buffer[end] = e
		r.size++
		return
	}

	old := r.buffer
	r.buffer = make([]interface{}, 2*len(old)+1)
	copy(r.buffer, old[r.off:])
	copy(r.buffer[len(old)-r.off:], old[:r.off])
	r.buffer[len(old)] = e

	r.off = 0
	r.size++
	return
}

// Pop the oldest element int the queue.
func (r *Ring) Pop() interface{} {
	if r.size == 0 {
		panic("Popped element from empty ring")
	}
	r.size--

	ret := r.buffer[r.off]
	r.buffer[r.off] = nil
	r.off = (r.off + 1) % len(r.buffer)
	return ret
}

type IntRing Ring

func (ir *IntRing) Add(e int) {
	(*Ring)(ir).Add(e)
}

// Pop the oldest element int the queue.
func (ir *IntRing) Pop() int {
	return (*Ring)(ir).Pop().(int)
}
