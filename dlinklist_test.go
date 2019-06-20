package dlinklist

import (
	"testing"
)

func TestDoublyLinkedList_PushFront(t *testing.T) {
	dll := DoublyLinkedList{}
	val := 1
	dll.PushFront(1)
	if val != dll.First().Value() {
		t.Errorf("expected that the first item has value %d, but he has %+v\n", val, dll.First().Value())
	}
}

func TestDoublyLinkedList_PushLast(t *testing.T) {
	dll := DoublyLinkedList{}
	val := 1
	dll.PushLast(1)
	if val != dll.Last().Value() {
		t.Errorf("expected that the last item has value %d, but he has %+v\n", val, dll.Last().Value())
	}
}

func TestDoublyLinkedList_Len(t *testing.T) {
	dll := DoublyLinkedList{}
	dll.PushFront(1)
	dll.PushFront(1)
	dll.PushLast(1)
	dll.PushLast(1)
	if dll.Len() != 4 {
		t.Errorf("expected that the list len is 4, %+v\n", dll)
	}
}

func TestDoublyLinkedList_First(t *testing.T) {
	dll := DoublyLinkedList{}
	val := 1
	dll.PushFront(1)
	if val != dll.First().Value() {
		t.Errorf("expected that the first item has value %d, but he has %+v\n", val, dll.First().Value())
		return
	}

	dll.First().Remove()
	if dll.First() != nil {
		t.Errorf("expected that the first item is nil, but it is %v", dll.First())
		return
	}
}

func TestDoublyLinkedList_Last(t *testing.T) {
	dll := DoublyLinkedList{}
	val := 1
	dll.PushLast(1)
	if val != dll.Last().Value() {
		t.Errorf("expected that the last item has value %d, but he has %+v\n", val, dll.Last().Value())
		return
	}

	dll.Last().Remove()
	if dll.Last() != nil {
		t.Errorf("expected that the last item is nil, but it is %v", dll.Last())
		return
	}
}

func TestItem_Remove(t *testing.T) {
	dll := DoublyLinkedList{}
	dll.PushFront(1)
	dll.First().Remove()

	if dll.Len() > 0 {
		t.Errorf("expected that the list is empty and an item was removed, %+v\n", dll)
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
