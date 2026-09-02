// Helper used by .github/workflows/build.yml to pack the module source zip.
//
// golang.org/x/mod/zip.CreateFromDir is Go's own module-zip writer, so it
// applies the module proxy's exact inclusion rules (dropping .git, vendor/ and
// nested modules) instead of a hand-rolled file list.
//
// It lives under .github/ for two reasons: the Go tool ignores directories
// whose name begins with a dot, so `go build ./...` and `go test ./...` never
// see it; and the packing step stages the module tree without .github, so it
// cannot leak into the published zip.
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: create_mod <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
