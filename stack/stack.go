package stack

type StackItem struct {
	item interface{}
	next *StackItem
}

type Stack struct {
	sp    *StackItem
	depth uint
}

func New() *Stack {
	var stack = new(Stack)
	stack.depth = 0
	return stack
}

func (stack *Stack) Push(data interface{}) {
	stack.sp = &StackItem{item: data, next: stack.sp}
	stack.depth++
}

func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty() {
		item := stack.sp.item
		stack.sp = stack.sp.next
		stack.depth--
		return item
	}
	return nil
}

func (stack *Stack) Peek() interface{} {
	if !stack.IsEmpty() {
		return stack.sp.item
	}
	return nil
}

func (stack *Stack) IsEmpty() bool {
	return stack.depth == 0
}
