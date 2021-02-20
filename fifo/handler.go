package fifo

import (
	"fmt"
	"reflect"
	"unsafe"
    "container/list"
)

// https://flaviocopes.com/golang-data-structure-queue/
/*Most queue implementations are in one of three flavors: slice-based, linked list-based, and circular-buffer (ring-buffer) based.

Slice-based queues tend to waste memory because they do not reuse the memory previously occupied by removed items. Also, slice based queues tend to only be single-ended.
Linked list queues can be better about memory reuse, but are generally a little slower and use more memory overall because of the overhead of maintaining links. They can offer the ability to add and remove items from the middle of the queue without moving memory around, but if you are doing much of that a queue is the wrong data structure.
Ring-buffer queues offer all the efficiency of slices, with the advantage of not wasting memory. Fewer allocations means better performance. They are just as efficient adding and removing items from either end so you naturally get a double-ended queue. So, as a general recommendation I would recommend a ring-buffer based queue implementation. This is what is discussed in the rest of this post.*/

// https://www.educative.io/edpresso/how-to-implement-a-queue-in-golang
func Handle() {
	version1()
	version2()
}

func version1() {
	var queue[] int // Make a queue of ints.

	queue = enqueue(queue, 10)
	queue = enqueue(queue, 20)
	queue = enqueue(queue, 30)

	fmt.Println("Queue:", queue)

	queue = dequeue(queue)
	fmt.Println("Queue:", queue)

	queue = enqueue(queue, 40)
	fmt.Println("Queue:", queue)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&queue))
	fmt.Println(*(*[3]int)(unsafe.Pointer(hdr.Data)))
}

func version2() {
	queue := list.New()

	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	front := queue.Front()
	fmt.Println(front.Value)
	queue.Remove(front)
}

func enqueue(queue[] int, element int) []int {
	queue = append(queue, element) // Simply append to enqueue.
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue[] int) ([]int) {
	element := queue[0] // The first element is the one to be dequeued.
	fmt.Println("Dequeued:", element)
	return queue[1:] // Slice off the element once it is dequeued.
}