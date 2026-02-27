package clang

import (
	"bufio"
	"fmt"
	"go/ast"
	"os"
	"strings"
)

// fail prints an error with source context to stderr and exits.
func (g *Generator) fail(node ast.Node, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	pos := g.pkg.Fset.Position(node.Pos())
	fmt.Fprintf(os.Stderr, "%s: %s\n", pos, msg)
	if srcLine, err := readSourceLine(pos.Filename, pos.Line); err == nil {
		fmt.Fprintf(os.Stderr, "%s\n", srcLine)
		fmt.Fprintf(os.Stderr, "%s^here\n", strings.Repeat("\t", pos.Column-1))
	}
	os.Exit(1)
}

// readSourceLine reads a single line from a source file (1-indexed).
func readSourceLine(filename string, line int) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for i := 1; scanner.Scan(); i++ {
		if i == line {
			return scanner.Text(), nil
		}
	}
	return "", fmt.Errorf("line %d not found in %s", line, filename)
}
