package render

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sourcery-zone/oza.sh/internal/helpers"
)

// TestRenderGitTemplate ...
func TestRenderGitTemplate(t *testing.T) {
	path := helpers.CreateTestGitRepo(t)

	render := func(format string) string {
		out, err := RenderStatus(format, path)
		if err != nil {
			t.Fatal(err)
		}
		return out
	}

	format := "{{ Git.Branch }}{{ if Git.Dirty }}*{{ end }}"
	if out := render(format); out != "main" {
		t.Errorf("unexpected output: %s", out)
	}

	// make it dirty
	os.WriteFile(filepath.Join(path, "TESTFILE"), []byte("# TEST"), 0644)
	if out := render(format); out != "main*" {
		t.Errorf("unexpected output: %s", out)
	}
}
