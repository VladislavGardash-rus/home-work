package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	firstItem *ListItem
	lastItem  *ListItem
}

func NewList() List {
	return new(list)
}

func (list *list) Len() int {
	length := 0

	item := list.firstItem
	if list.firstItem != nil {
		for {
			length++

			if item.Next == nil {
				break
			}

			item = item.Next
		}
	}

	return length
}

func (list *list) Front() *ListItem {
	return list.firstItem
}

func (list *list) Back() *ListItem {
	return list.lastItem
}

func (list *list) PushFront(value interface{}) *ListItem {
	if list.firstItem == nil {
		insertFirstItem(list, value)
	} else {
		firstItem := list.firstItem
		item := &ListItem{
			Value: value,
			Prev:  nil,
			Next:  firstItem,
		}

		firstItem.Prev = item
		list.firstItem = item
	}

	return list.firstItem
}

func (list *list) PushBack(value interface{}) *ListItem {
	if list.lastItem == nil {
		insertFirstItem(list, value)
	} else {
		item := &ListItem{
			Value: value,
			Prev:  list.lastItem,
			Next:  nil,
		}

		list.lastItem.Next = item
		list.lastItem = item
	}

	return list.lastItem
}

func (list *list) Remove(item *ListItem) {
	if item.Prev == nil {
		list.firstItem = nil
		if item.Next != nil {
			list.firstItem = item.Next
			item.Next.Prev = nil
		}
	}

	if item.Next == nil {
		list.lastItem = nil
		if item.Prev != nil {
			list.lastItem = item.Prev
			item.Prev.Next = nil
		}
	}

	if item.Prev != nil {
		if item.Next != nil {
			item.Next.Prev = item.Prev
		}
	}

	if item.Next != nil {
		if item.Prev != nil {
			item.Prev.Next = item.Next
		}
	}
}

func (list *list) MoveToFront(item *ListItem) {
	list.Remove(item)
	item.Prev = nil
	item.Next = list.firstItem
	list.firstItem.Prev = item
	list.firstItem = item
}

func insertFirstItem(list *list, value interface{}) {
	item := &ListItem{
		Value: value,
		Prev:  nil,
		Next:  nil,
	}

	list.firstItem = item
	list.lastItem = item
}
