package memory

type MEMORY_TYPE byte
type MEMORY []MEMORY_TYPE

var (
	memory     MEMORY
	memorySize uint
	pointer    uint
)

func Init(size uint) {
	memorySize = size
	memory = make([]MEMORY_TYPE, size)
}

func PointerMoveLeft() {
	pointer = (pointer + memorySize - 1) % memorySize
}

func PointerMoveRight() {
	pointer = (pointer + memorySize + 1) % memorySize
}

func GetPointer() uint {
	return pointer
}

func Get() MEMORY_TYPE {
	return MEMORY_TYPE(memory[pointer])
}

func Set(value MEMORY_TYPE) {
	memory[pointer] = MEMORY_TYPE(value)
}

func Incr() {
	memory[pointer]++
}

func Decr() {
	memory[pointer]--
}
