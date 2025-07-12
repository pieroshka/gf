# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is "gf" - a Brainfuck to Go transpiler. The project takes Brainfuck source code and transpiles it into Go code that can be executed natively or compiled to WebAssembly.

## Architecture

The transpiler follows a three-stage pipeline:

1. **Lexical Analysis** (`internal/analysis/lexical.go`) - Filters out non-Brainfuck tokens from source
2. **Syntactical Analysis** (`internal/analysis/syntactical.go`) - Validates bracket matching 
3. **Code Generation** (`internal/compile/compile.go`) - Transpiles to Go using github.com/dave/jennifer

### Key Components

- `cmd/gf/main.go` - CLI entry point with file I/O handling
- `internal/token/token.go` - Defines the 8 Brainfuck tokens (`<>+-.,[]`)
- `internal/analysis/` - Two-stage analysis pipeline
- `internal/compile/` - Go code generation with consecutive operation optimization

The compiler optimizes consecutive identical operations (e.g., `+++` becomes `memory[pointer] += 3`) and generates Go code with proper imports and error handling.

## Development Commands

Build and run the transpiler:
```bash
# Transpile and run (default - no WebAssembly)
task

# Transpile and compile to WebAssembly
task wasm

# Run with custom input
go run cmd/gf/main.go -in <brainfuck_file> -out <output_go_file>
```

Testing:
```bash
# Run all tests
go test ./...

# Run specific test package
go test ./internal/analysis
```

Build operations:
```bash
# Build the CLI
go build -o gf cmd/gf/main.go

# Module operations
go mod tidy
go mod download
```

## Sample Brainfuck Files

The `bf/` directory contains test Brainfuck programs:
- `hello_world.bf` - Classic "Hello World!" program with detailed comments
- `bad1.bf`, `bad2.bf` - Invalid programs for testing error handling