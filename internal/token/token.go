package token

var (
	PtrMoveLeft   byte = '<'
	PtrMoveRight  byte = '>'
	IncrMemCell   byte = '+'
	DecrMemCell   byte = '-'
	OutputMemCell byte = '.'
	InputMemCell  byte = ','
	BracketOpen   byte = '['
	BracketClose  byte = ']'

	Tokens = []byte{
		PtrMoveLeft,
		PtrMoveRight,
		IncrMemCell,
		DecrMemCell,
		OutputMemCell,
		InputMemCell,
		BracketOpen,
		BracketClose,
	}
)
