package memory

type memoryNode struct {
	prev *memoryNode
	next *memoryNode

	data byte
}

var (
	pointer *memoryNode
)

func Init() {
	indexZero := memoryNode{}
	pointer = &indexZero
}

func PointerMoveLeft() {
	if pointer.prev == nil {
		pointer.prev = &memoryNode{next: pointer}
	}
	pointer = pointer.prev
}

func PointerMoveRight() {
	if pointer.next == nil {
		pointer.next = &memoryNode{prev: pointer}
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
