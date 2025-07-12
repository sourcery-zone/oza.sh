package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sourcery-zone/oza.sh/internal/render"
)

// main oza.sh: fast and extensible command-line to build feature-rich status lines!
func main()  {
	if len(os.Args) < 3 {
		fmt.Println("Usage: oza <path> <format>")
		os.Exit(1)
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid path: %v", err)
	}

	format := os.Args[2]
	out, err := render.RenderStatus(format, path)
	if err != nil {
		log.Fatalf("Render error: %v", err)
	}

	fmt.Println(out)
}
