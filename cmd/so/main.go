package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/nalgeon/soan/internal/compiler"
)

func main() {
	var dir string
	var outDir string
	flag.StringVar(&dir, "dir", "", "source directory (.go)")
	flag.StringVar(&outDir, "out", "", "output directory (.c)")
	flag.Parse()

	if dir == "" {
		slog.Error("must specify -dir")
		os.Exit(1)
	}
	if outDir == "" {
		outDir = dir
	}

	// Set up logging
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	compiler.Translate(dir, outDir)
}
