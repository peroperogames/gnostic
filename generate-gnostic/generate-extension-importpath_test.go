// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// writeModule writes a go.mod with the specified contents into dir and returns dir.
func writeModule(t *testing.T, dir, contents string) string {
	t.Helper()
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("cannot create %s: %v", dir, err)
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(contents), 0644); err != nil {
		t.Fatalf("cannot write go.mod in %s: %v", dir, err)
	}
	return dir
}

func TestModuleForDirectory(t *testing.T) {
	for _, tt := range []struct {
		name     string
		goMod    string
		subDir   string
		wantPath string
		wantOK   bool
	}{
		{
			name:     "simple module path",
			goMod:    "module example.com/simple\n\ngo 1.24.6\n",
			wantPath: "example.com/simple",
			wantOK:   true,
		},
		{
			name:     "quoted module path is unquoted",
			goMod:    "module \"example.com/quoted\"\n\ngo 1.24.6\n",
			wantPath: "example.com/quoted",
			wantOK:   true,
		},
		{
			name:     "module path with a trailing comment",
			goMod:    "module example.com/commented // the module path\n\ngo 1.24.6\n",
			wantPath: "example.com/commented",
			wantOK:   true,
		},
		{
			name:     "module path separated by a tab",
			goMod:    "module\texample.com/tabbed\n\ngo 1.24.6\n",
			wantPath: "example.com/tabbed",
			wantOK:   true,
		},
		{
			name:     "walk up to the module root",
			goMod:    "module example.com/nested\n\ngo 1.24.6\n",
			subDir:   "extensions/sample/generated/gnostic-x-demo/proto",
			wantPath: "example.com/nested",
			wantOK:   true,
		},
		{
			name:   "go.mod without a module directive",
			goMod:  "go 1.24.6\n",
			wantOK: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			root := writeModule(t, t.TempDir(), tt.goMod)
			dir := root
			if tt.subDir != "" {
				dir = filepath.Join(root, filepath.FromSlash(tt.subDir))
				if err := os.MkdirAll(dir, 0755); err != nil {
					t.Fatalf("cannot create %s: %v", dir, err)
				}
			}

			gotRoot, gotPath, ok := moduleForDirectory(dir)
			if ok != tt.wantOK {
				t.Fatalf("moduleForDirectory(%q) ok = %v, want %v", dir, ok, tt.wantOK)
			}
			if !tt.wantOK {
				return
			}
			if gotPath != tt.wantPath {
				t.Errorf("module path = %q, want %q", gotPath, tt.wantPath)
			}

			// The returned root has to be the directory holding the go.mod, which
			// the caller relies on to make the output directory relative.
			wantRoot, err := filepath.EvalSymlinks(root)
			if err != nil {
				t.Fatalf("cannot resolve %s: %v", root, err)
			}
			resolvedRoot, err := filepath.EvalSymlinks(gotRoot)
			if err != nil {
				t.Fatalf("cannot resolve %s: %v", gotRoot, err)
			}
			if resolvedRoot != wantRoot {
				t.Errorf("module root = %q, want %q", resolvedRoot, wantRoot)
			}
		})
	}
}

// importLines returns the entries of the import block of a generated Go file.
func importLines(t *testing.T, contents string) []string {
	t.Helper()
	lines := []string{}
	inBlock := false
	for _, line := range strings.Split(contents, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "import (" {
			inBlock = true
			continue
		}
		if inBlock && trimmed == ")" {
			inBlock = false
			continue
		}
		if inBlock {
			lines = append(lines, trimmed)
		}
	}
	if len(lines) == 0 {
		t.Fatalf("no import block found in:\n%s", contents)
	}
	return lines
}

// TestGeneratedImportPaths runs the generator against a throwaway module and
// checks the import paths written into the generated handler.
//
// They have to follow the module that contains the output directory regardless
// of how that module is named and of where the tool is run from: building them
// from a fixed "extensions/sample" location produced unbuildable paths for any
// other output directory, and taking a quoted module path verbatim produced
// invalid Go.
func TestGeneratedImportPaths(t *testing.T) {
	if _, err := exec.LookPath("generate-gnostic"); err != nil {
		t.Skip("generate-gnostic is not on PATH: install it with 'go install ./generate-gnostic'")
	}

	root := writeModule(t, t.TempDir(), "module \"example.com/quoted\"\n\ngo 1.24.6\n")

	// Deliberately not the sample layout, and not the directory the tool is run
	// from, so a hardcoded layout or a working-directory assumption would show.
	outDir := filepath.Join(root, "somewhere", "else")
	schema, err := filepath.Abs(filepath.Join("..", "extensions", "sample", "x-sampleone.json"))
	if err != nil {
		t.Fatalf("cannot resolve the sample schema: %v", err)
	}

	command := exec.Command("generate-gnostic", "--extension", schema, "--out_dir="+outDir)
	command.Dir = root
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("generate-gnostic failed: %v\n%s", err, output)
	}

	generated, err := os.ReadFile(filepath.Join(outDir, "gnostic-x-sampleone", "main.go"))
	if err != nil {
		t.Fatalf("cannot read the generated handler: %v", err)
	}
	contents := string(generated)
	imports := importLines(t, contents)

	for _, want := range []string{
		`"example.com/quoted/extensions"`,
		`"example.com/quoted/compiler"`,
		`"example.com/quoted/somewhere/else/gnostic-x-sampleone/proto"`,
	} {
		found := false
		for _, line := range imports {
			if line == want {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("generated imports %v are missing %s", imports, want)
		}
	}

	// The import path used to be built by concatenating a fixed location onto an
	// arbitrary output directory, which glued an absolute path on and left a
	// doubled slash behind.
	for _, unwanted := range []string{"extensions/sample", "//", `"/`} {
		for _, line := range imports {
			if strings.Contains(line, unwanted) {
				t.Errorf("generated import %s contains %q, so the path is malformed", line, unwanted)
			}
		}
	}
}
