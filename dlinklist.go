// Package dlinklist implements a doubly linked list structure.
package dlinklist

// DoublyLinkedList represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type DoublyLinkedList struct {
	len   int
	first *Item
	last  *Item
}

// Len returns the number of items of the list dll.
// The complexity is O(1).
func (dll *DoublyLinkedList) Len() int {
	return dll.len
}

// First returns the first item of list dll (or nil if the list is empty).
func (dll *DoublyLinkedList) First() *Item {
	return dll.first
}

// Last returns the last item of list dll (or nil if the list is empty).
func (dll *DoublyLinkedList) Last() *Item {
	return dll.last
}

// PushFront pushes a new item with value val at the front of list dll.
func (dll *DoublyLinkedList) PushFront(val interface{}) {
	newItem := &Item{value: val, next: dll.first}
	dll.len++
	newItem.list = dll
	// If the list length equals 1 then the new Item is the only in the list,
	// so the item is the first and the last.
	if dll.len == 1 {
		dll.first = newItem
		dll.last = newItem
		return
	}

	dll.first.prev = newItem
	dll.first = newItem
}

// PushLast pushes a new item with value val at the end of list dll.
func (dll *DoublyLinkedList) PushLast(val interface{}) {
	newItem := &Item{value: val, prev: dll.last}
	dll.len++
	newItem.list = dll
	// If the list length equals 1 then the new Item is the only in the list,
	// so the item is the first and the last.
	if dll.Len() == 1 {
		dll.first = newItem
		dll.last = newItem
		return
	}

	dll.last.next = newItem
	dll.last = newItem
}

// Item is an item of a doubly linked list.
type Item struct {
	// Value of the item.
	value interface{}

	// Pointers in the doubly linked list.
	next, prev *Item

	// A doubly linked list to which this item belongs.
	list *DoublyLinkedList
}

// Value returns value of the item i.
func (i *Item) Value() interface{} {
	return i.value
}

// Next returns the next list item or nil.
func (i *Item) Next() *Item {
	return i.next
}

// Prev returns the previous list item or nil.
func (i *Item) Prev() *Item {
	return i.prev
}

// Remove removes the item from its list and decrements list.len.
func (i *Item) Remove() {
	i.list.len--
	// If it was the last item in the list, we just reset the list pointers.
	if i.list.len == 0 {
		i.list.first = nil
		i.list.last = nil
		return
	}

	if i.next != nil {
		i.next.prev = i.prev

		// If the deleting item is the first, we should move the list first pointer to the next item.
		if i == i.list.first {
			i.list.first = i.next
		}
	}

	if i.prev != nil {
		i.prev.next = i.next

		// If the deleting item is the last, we should move the list last pointer to the prev item.
		if i == i.list.last {
			i.list.last = i.prev
		}
	}
}
