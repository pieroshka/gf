# gf - Brainfuck to Go Transpiler

A high-performance Brainfuck to Go transpiler that converts Brainfuck programs into optimized Go code with support for both native execution and WebAssembly compilation.

## Features

- **Fast Transpilation**: Three-stage pipeline (lexical → syntactical → code generation)
- **Optimization**: Consecutive operations are optimized (e.g., `+++` becomes `memory[pointer] += 3`)
- **WebAssembly Support**: Compile transpiled Go code to WASM using TinyGo
- **Error Handling**: Comprehensive syntax validation with precise error reporting
- **Clean Output**: Generated Go code is properly formatted and commented

## Installation

```bash
git clone https://github.com/pieroshka/gf.git
cd gf
go mod download
```

## Usage

### Basic Usage

```bash
# Transpile and run a Brainfuck program
go run cmd/gf/main.go -in bf/hello_world.bf

# Specify custom output file
go run cmd/gf/main.go -in program.bf -out output.go

# Read from stdin
echo "+++++++++[>++++++++++<-]>+++++++++++.>" | go run cmd/gf/main.go -in -
```

### Using Task (Recommended)

```bash
# Transpile and run (native Go)
task

# Transpile and compile to WebAssembly
task wasm
```

### Building

```bash
# Build the transpiler
go build -o gf cmd/gf/main.go

# Run built binary
./gf -in bf/hello_world.bf
```

## How It Works

The transpiler processes Brainfuck code through three stages:

1. **Lexical Analysis**: Filters out non-Brainfuck characters, keeping only the 8 valid tokens: `<>+-.,[]`
2. **Syntactical Analysis**: Validates bracket matching and reports syntax errors with precise locations
3. **Code Generation**: Transpiles to optimized Go code using the [Jennifer](https://github.com/dave/jennifer) code generator

### Optimization

The transpiler automatically optimizes consecutive identical operations:

```brainfuck
+++++ → memory[pointer] += 5
----- → memory[pointer] -= 5
>>>>> → pointer += 5
<<<<< → pointer -= 5
```

## Examples

### Hello World

```brainfuck
+++++ +++               Set Cell #0 to 8
[
    >++++               Add 4 to Cell #1
    [
        >++             Add 4*2 to Cell #2
        >+++            Add 4*3 to Cell #3
        >+++            Add 4*3 to Cell #4
        >+              Add 4 to Cell #5
        <<<<-           Decrement the loop counter in Cell #1
    ]
    >+                  Add 1 to Cell #2
    >+                  Add 1 to Cell #3
    >-                  Subtract 1 from Cell #4
    >>+                 Add 1 to Cell #6
    [<]                 Move back to the first zero cell
    <-                  Decrement the loop Counter in Cell #0
]
>>.                     Output 'H'
>---.                   Output 'e'
+++++ ++..+++.          Output 'llo'
>>.                     Output ' '
<-.                     Output 'W'
<.                      Output 'o'
+++.----- -.----- ---.  Output 'rld'
>>+.                    Output '!'
>++.                    Output newline
```

### Generated Go Code Structure

```go
package main

func main() {
    memory := make([]int, 30000)
    var pointer int
    
    memory[pointer] += 8     // ++++++++
    for memory[pointer] != 0 {
        pointer += 1         // >
        memory[pointer] += 4 // ++++
        // ... optimized operations
    }
    // ... rest of the program
}
```

## Architecture

```
cmd/gf/           # CLI entry point
internal/
├── token/        # Brainfuck token definitions
├── analysis/     # Lexical and syntactical analysis
│   ├── lexical.go
│   └── syntactical.go
└── compile/      # Go code generation
    └── compile.go
```

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./internal/analysis
go test ./internal/compile

# Run with verbose output
go test -v ./...
```

### Testing with Sample Programs

The `bf/` directory contains sample Brainfuck programs:
- `hello_world.bf` - Classic "Hello World!" with detailed comments
- `bad1.bf`, `bad2.bf` - Invalid programs for testing error handling

## Dependencies

- [Jennifer](https://github.com/dave/jennifer) - Go code generation
- [Testify](https://github.com/stretchr/testify) - Testing utilities
- [Task](https://taskfile.dev/) - Task runner (optional)
- [TinyGo](https://tinygo.org/) - WebAssembly compilation (optional)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the MIT License.

## Brainfuck Language Reference

| Command | Description |
|---------|-------------|
| `>`     | Move pointer to the right |
| `<`     | Move pointer to the left |
| `+`     | Increment memory cell at pointer |
| `-`     | Decrement memory cell at pointer |
| `.`     | Output character at memory cell |
| `,`     | Input character to memory cell |
| `[`     | Jump past matching `]` if cell is 0 |
| `]`     | Jump back to matching `[` if cell is not 0 |