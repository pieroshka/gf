package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
)

var Token = struct {
	PtrMoveLeft   byte
	PtrMoveRight  byte
	IncrMemCell   byte
	DecrMemCell   byte
	OutputMemCell byte
	InputMemCell  byte
	BracketOpen   byte
	BracketClose  byte
}{
	PtrMoveLeft:   '<',
	PtrMoveRight:  '>',
	IncrMemCell:   '+',
	DecrMemCell:   '-',
	OutputMemCell: '.',
	InputMemCell:  ',',
	BracketOpen:   '[',
	BracketClose:  ']',
}

var (
	TOKENS = []byte{
		Token.PtrMoveLeft,
		Token.PtrMoveRight,
		Token.IncrMemCell,
		Token.DecrMemCell,
		Token.OutputMemCell,
		Token.InputMemCell,
		Token.BracketOpen,
		Token.BracketClose,
	}
)

func isToken(b byte) bool {
	for i := range TOKENS {
		if b == TOKENS[i] {
			return true
		}
	}
	return false
}

func lexical(bfcode []byte) []byte {
	var out []byte
	for i := 0; i < len(bfcode); i++ {
		if isToken(bfcode[i]) {
			out = append(out, bfcode[i])
		}
	}
	return out
}

func syntactical(bfcode []byte) (int, bool) {
	var stack []byte

	for i := range bfcode {
		if bfcode[i] == Token.BracketOpen {
			stack = append(stack, Token.BracketOpen)
		}

		if bfcode[i] == Token.BracketClose {
			if len(stack) < 1 || stack[len(stack)-1] != Token.BracketOpen {
				return i, false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(bfcode), len(stack) == 0
}

type executable struct{}

func translateBFRec(source []byte, ptr *int) []jen.Code {
	statements := []jen.Code{}

loop:
	for ; *ptr < len(source); *ptr++ {
		switch source[*ptr] {
		case Token.PtrMoveLeft:
			statements = append(statements, jen.Id("pointer").Op("--").Comment(string(source[*ptr])))
		case Token.PtrMoveRight:
			statements = append(statements, jen.Id("pointer").Op("++").Comment(string(source[*ptr])))
		case Token.IncrMemCell:
			statements = append(statements, jen.Id("memory").Index(jen.Id("pointer")).Op("++").Comment(string(source[*ptr])))
		case Token.DecrMemCell:
			statements = append(statements, jen.Id("memory").Index(jen.Id("pointer")).Op("--").Comment(string(source[*ptr])))
		case Token.OutputMemCell:
			statements = append(statements, jen.Qual("fmt", "Printf").Call(jen.Lit("%c"), jen.Id("memory").Index(jen.Id("pointer"))).Comment(string(source[*ptr])))
		case Token.InputMemCell:
			statements = append(statements, jen.Id("reader").Op("=").Qual("bufio", "NewReader").Call(jen.Qual("os", "Stdin")).Comment(string(source[*ptr])))
			statements = append(statements, jen.List(jen.Id("c"), jen.Id("_"), jen.Id("err").Op("=").Id("reader").Dot("ReadRune").Call()))
			statements = append(statements, jen.If(jen.Id("err").Op("!=").Nil().Block(
				jen.Panic(jen.Id("err")),
			)))
			statements = append(statements, jen.Id("memory").Index(jen.Id("pointer")).Op("=").Int().Parens(jen.Id("c")))
		case Token.BracketOpen:
			*ptr++
			statements = append(statements, jen.For((jen.Id("memory").Index(jen.Id("pointer")).Op("!=").Lit(0)).Block(translateBFRec(source, ptr)...).Comment("]")))
		case Token.BracketClose:
			break loop
		}
	}

	return statements
}

func translateBF(source []byte) []jen.Code {
	ptr := 0
	return translateBFRec(source, &ptr)
}

func compile(source []byte) (*executable, error) {
	f := jen.NewFile("main")
	statements := []jen.Code{
		jen.Id("memory").Op(":=").Make(jen.Index().Int(), jen.Lit(30000)),
		jen.Var().Id("pointer").Int(),
	}

	if strings.Contains(string(source), ",") {
		// otherwise those vars are unused
		statements = append(statements,
			jen.Var().Id("reader").Op("*").Qual("bufio", "Reader"),
			jen.Var().Id("c").Rune(),
			jen.Var().Id("err").Error(),
		)
	}

	statements = append(statements, jen.Empty())
	statements = append(statements, translateBF(source)...)
	f.Func().Id("main").Params().Block(statements...)

	buf := new(bytes.Buffer)
	err := f.Render(buf)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll("out/", 0755)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile("out/main.go", buf.Bytes(), 0755)
	if err != nil {
		return nil, err
	}

	fmt.Println(buf.String())
	return new(executable), nil
}

func main() {
	bfcode := []byte("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.") // hello world
	// bfcode := []byte("+[[->]-[-<]>-]>.>>>>.<<<<-.>>-.>.<<.>>>>-.<<<<<++.>>++.")
	// bfcode := []byte(",.,.,.")
	// bfcode := []byte("+[++[+++]++++]+++++")
	bfcode = lexical(bfcode)
	i, ok := syntactical(bfcode)
	if !ok {
		log.Fatalln("syntax error at position:", i)
	}

	_, err := compile(bfcode)
	if err != nil {
		fmt.Println(err)
	}
}
