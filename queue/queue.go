package queue

type QueueItem struct {
	item interface{}
	next *QueueItem
}

type Queue struct {
	first *QueueItem
	last  *QueueItem
	len   uint
}

func New() *Queue {
	newQueue := new(Queue)
	newQueue.len = 0
	return newQueue
}

// Will append as a last item to the queue
func (queue *Queue) Enqueue(item interface{}) *Queue {
	queueItem := &QueueItem{item: item, next: nil}

	if queue.IsEmpty() {
		queue.first = queueItem
		queue.last = queueItem
		queue.len++
		return queue
	}

	queue.last.next = queueItem
	queue.last = queue.last.next
	queue.len++

	return queue
}

// Will remove the item in FIFO fashion
func (queue *Queue) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	}

	firstItem := queue.first
	queue.first = firstItem.next
	queue.len--

	return firstItem.item
}

func (queue *Queue) Peek() *QueueItem {
	if !queue.IsEmpty() {
		return queue.first
	}
	return nil
}

func (queue *Queue) IsEmpty() bool {
	return queue.len == 0
}
