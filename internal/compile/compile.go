package compile

import (
	"bytes"
	"os"
	"path"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/pieroshka/gf/internal/token"
)

type compiler struct{}

func New() *compiler {
	return &compiler{}
}

func (c compiler) Run(source []byte, outPath string) error {
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
	statements = append(statements, transpile(source)...)
	f.Func().Id("main").Params().Block(statements...)

	buf := new(bytes.Buffer)
	err := f.Render(buf)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Dir(outPath), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(outPath, buf.Bytes(), 0755)
	if err != nil {
		return err
	}

	return nil
}

func transpileRec(source []byte, ptr *int) []jen.Code {
	statements := []jen.Code{}

loop:
	for ; *ptr < len(source); *ptr++ {
		consecutive := 1
		for *ptr < len(source)-1 &&
			source[*ptr] == source[*ptr+1] &&
			(source[*ptr] == token.PtrMoveLeft ||
				source[*ptr] == token.PtrMoveRight ||
				source[*ptr] == token.IncrMemCell ||
				source[*ptr] == token.DecrMemCell) {
			consecutive++
			*ptr++
		}

		comment := strings.Repeat(string(source[*ptr]), consecutive)

		switch source[*ptr] {
		case token.PtrMoveLeft:
			statements = append(statements, jen.Id("pointer").Op("-=").Lit(consecutive).Comment(comment))
		case token.PtrMoveRight:
			statements = append(statements, jen.Id("pointer").Op("+=").Lit(consecutive).Comment(comment))
		case token.IncrMemCell:
			statements = append(statements, jen.Id("memory").Index(jen.Id("pointer")).Op("+=").Lit(consecutive).Comment(comment))
		case token.DecrMemCell:
			statements = append(statements, jen.Id("memory").Index(jen.Id("pointer")).Op("-=").Lit(consecutive).Comment(comment))
		case token.OutputMemCell:
			statements = append(statements, jen.Qual("fmt", "Printf").Call(jen.Lit("%c"), jen.Id("memory").Index(jen.Id("pointer"))).Comment(comment))
		case token.InputMemCell:
			statements = append(statements, jen.Id("reader").Op("=").Qual("bufio", "NewReader").Call(jen.Qual("os", "Stdin")).Comment(comment))
			statements = append(statements, jen.List(jen.Id("c"), jen.Id("_"), jen.Id("err").Op("=").Id("reader").Dot("ReadRune").Call()))
			statements = append(statements, jen.If(jen.Id("err").Op("!=").Nil().Block(
				jen.Panic(jen.Id("err")),
			)))
			statements = append(statements, jen.Id("memory").Index(jen.Id("pointer")).Op("=").Int().Parens(jen.Id("c")))
		case token.BracketOpen:
			*ptr++
			statements = append(statements, jen.For((jen.Id("memory").Index(jen.Id("pointer")).Op("!=").Lit(0)).Block(transpileRec(source, ptr)...).Comment(comment)))
		case token.BracketClose:
			break loop
		}
	}

	return statements
}

func transpile(source []byte) []jen.Code {
	ptr := 0
	return transpileRec(source, &ptr)
}
