package compiler

import (
	"embed"
	"log/slog"
	"os"
	"path/filepath"
)

//go:embed runtime/so.h runtime/so.c
var runtimeFS embed.FS

func writeRuntime(outDir string) {
	for _, name := range []string{"so.h", "so.c"} {
		data, err := runtimeFS.ReadFile("runtime/" + name)
		if err != nil {
			slog.Error("failed to read embedded runtime file", "name", name, "error", err)
			os.Exit(1)
		}
		if err := os.WriteFile(filepath.Join(outDir, name), data, 0o644); err != nil {
			slog.Error("failed to write runtime file", "error", err)
			os.Exit(1)
		}
	}
}
