package CircularQueue

// SimpleRingBuffer Reference: [Golang program to implement a circular buffer](https://www.tutorialspoint.com/golang-program-to-implement-a-circular-buffer)
type SimpleRingBuffer struct {
	buffer     []int
	size       int
	writeIndex int
	readIndex  int
	count      int
}

func constructor(k int) SimpleRingBuffer {
	return SimpleRingBuffer{
		buffer: make([]int, k),
		size:   k,
	}
}

func (s *SimpleRingBuffer) EnQueue(value int) bool {
	if s.count == s.size {
		s.readIndex = (s.readIndex + 1) % s.size
	} else {
		s.count++
	}
	s.buffer[s.writeIndex] = value
	s.writeIndex = (s.writeIndex + 1) % s.size
	return true
}

func (s *SimpleRingBuffer) Front() int {
	return s.buffer[s.readIndex]
}

func (s *SimpleRingBuffer) Rear() int {
	return s.buffer[s.writeIndex]
}
