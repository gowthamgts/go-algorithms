package queue

type QueueItem struct {
	item interface{}
	next *QueueItem
}

type Queue struct {
	first   *QueueItem
	current *QueueItem
	len     uint
}

func New() *Queue {
	newQueue := new(Queue)
	newQueue.len = 0
	return newQueue
}

func (queue *Queue) Push(item interface{}) *Queue {
	queueItem := &QueueItem{item: item, next: nil}

	if queue.len == 0 {
		// insert as first element
		queue.first = queueItem
		queue.current = queueItem
		queue.len++
		return queue
	}
	
	queue.current.next = queueItem
	queue.current = queueItem
	queue.len++
	return queue
}

func (queue *Queue) Pop() *QueueItem {
	if queue.len > 0 {
		popItem := queue.first
		queue.first = popItem.next
		queue.len--
		return popItem
	}
	return nil
}

func (queue *Queue) Peek() *QueueItem {
	if queue.len > 0 {
		return queue.first
	}
	return nil
}
