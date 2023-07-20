package memory

type memoryDLL struct {
	head *memoryNode
	tail *memoryNode
}

type memoryNode struct {
	prev *memoryNode
	next *memoryNode

	data byte
}

var (
	memory  memoryDLL
	pointer *memoryNode
)

func Init() {
	indexZero := memoryNode{}
	memory = memoryDLL{head: &indexZero, tail: &indexZero}
	pointer = memory.head
}

func PointerMoveLeft() {
	if pointer.prev == nil {
		pointer.prev = &memoryNode{next: pointer}
		memory.head = pointer.prev
	}
	pointer = pointer.prev
}

func PointerMoveRight() {
	if pointer.next == nil {
		pointer.next = &memoryNode{prev: pointer}
		memory.tail = pointer.next
	}
	pointer = pointer.next
}

func Get() byte {
	return pointer.data
}

func Set(value byte) {
	pointer.data = value
}

func Decr() {
	pointer.data--
}

func Incr() {
	pointer.data++
}
