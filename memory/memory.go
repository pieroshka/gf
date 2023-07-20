package memory

type MEMORY []byte

var (
	memory     MEMORY
	memorySize uint
	pointer    uint
)

func Init(size uint) {
	memorySize = size
	memory = make([]byte, size)
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

func Get() byte {
	return memory[pointer]
}

func Set(value byte) {
	memory[pointer] = value
}

func Incr() {
	memory[pointer]++
}

func Decr() {
	memory[pointer]--
}
