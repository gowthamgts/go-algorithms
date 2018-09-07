package list

type ListItem struct {
	prev *ListItem
	next *ListItem
	item interface{}
}

type List struct {
	first *ListItem
	last  *ListItem
	depth uint
}

func New() *List {
	list := new(List)
	list.depth = 0
	return list
}

func (list *List) isEmpty() bool {
	return list.depth == 0
}


func (list *List) Prepend(item interface{}) *List {
	if list.isEmpty() {
		return list.InsertToEmptyList(item)
	}

	listItem := &ListItem{prev: nil, next: list.first, item: item}
	list.first.prev = listItem
	list.first = listItem
	list.depth++

	return list
}

func (list *List) Append(item interface{}) *List {
	if list.isEmpty() {
		return list.InsertToEmptyList(item)
	}

	listItem := &ListItem{prev: list.last, next: nil, item: item}
	list.last.next = listItem
	list.last = listItem
	list.depth++

	return list
}

func (list *List) Get(position uint) *ListItem {
	if list.isEmpty() {
		return nil
	}

	listItem := list.first

	for i := uint(0); i < position; i++ {
		listItem = listItem.next
	}

	return listItem
}

func (list *List) InsertToEmptyList(item interface{}) *List {
	listItem := &ListItem{prev: nil, next: nil, item: item}
	list.first = listItem
	list.last = listItem
	list.depth++
	return list
}
