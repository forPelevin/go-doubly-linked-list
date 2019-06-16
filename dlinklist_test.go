package dlinklist

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	// Initialize a new doubly linked list
	dll := DoublyLinkedList{}

	// Push a new item at the beginning of the list.
	dll.PushFront(1)

	// The new item should be the first and the last node.
	firstItem := dll.First()
	lastItem := dll.Last()
	if firstItem == nil || lastItem == nil || firstItem.Value() != lastItem.Value() {
		t.Errorf(
			"expected that the first and last item are same, but they aren't,\n%+v\n%+v",
			firstItem,
			lastItem,
		)
		return
	}

	// And the list length must be 1
	if dll.Len() != 1 {
		t.Errorf("expected that the list len is 1, %+v\n", dll)
		return
	}

	// Push a new item at the beginning again.
	val2 := 2
	dll.PushFront(2)

	// Now the new item must be first.
	if val2 != dll.First().Value() {
		t.Errorf("expected that the first item has value %d, but he has %+v\n", val2, dll.First().Value())
		return
	}

	// And the last item must point to the new item as prev,
	if dll.Last().Prev() != dll.First() {
		t.Errorf("expected that the last item prev is the new item %+v, but got %+v\n", val2, dll.First().Value())
		return
	}

	// but still has next as nil.
	if dll.Last().Next() != nil {
		t.Errorf("expected that the last item next is nil, but got %+v\n", dll.Last().Next())
		return
	}

	// Push a new item at the end of the list.
	val3 := 3
	dll.PushLast(val3)

	// The new item must be the last.
	if val3 != dll.Last().Value() {
		t.Errorf("expected that the last item has value %d, but he has %+v\n", val3, dll.Last().Value())
		return
	}

	// And the list length should be 3.
	if dll.Len() != 3 {
		t.Errorf("expected that the list len is 3, %+v\n", dll)
		return
	}

	// Remove the last and middle items.
	dll.Last().Remove()
	dll.Last().Remove()

	// The list len must be 1.
	if dll.Len() != 1 {
		t.Errorf("expected that the list len is 1, %+v\n", dll)
		return
	}

	// The list first and last items must be one item.
	if dll.First() != dll.Last() {
		t.Errorf(
			"expected that the first and last item are same, but they aren't,\n%+v\n%+v",
			dll.First(),
			dll.Last(),
		)
		return
	}

	// The remaining must have empty prev and last pointers.
	remainingItem := dll.First()
	if remainingItem.next != nil || remainingItem.prev != nil {
		t.Errorf("expected that the remaining item has nil prev and next pointers, but got %+v\n", remainingItem)
		return
	}

	// Delete the remaining item.
	remainingItem.Remove()

	// The list len must be 0.
	if dll.Len() != 0 {
		t.Errorf("expected that the list len is 0, %+v\n", dll)
		return
	}

	// The list first and last items must be nils.
	if dll.First() != nil || dll.Last() != nil {
		t.Errorf(
			"expected that the first and last item are nils, but they aren't,\n%+v\n%+v",
			dll.First(),
			dll.Last(),
		)
		return
	}
}

func BenchmarkDoublyLinkedList(b *testing.B) {
	dll := DoublyLinkedList{}
	for i := 0; i < b.N; i++ {
		dll.PushFront(1)
		dll.PushLast(1)
		dll.Len()

		if i%10 == 0 {
			dll.First().Remove()
			dll.Last().Remove()
		}
	}
}
