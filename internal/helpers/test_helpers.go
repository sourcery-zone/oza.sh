package helpers

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// runCmd run a command line in a directory and fail if error happened.
func RunCmd(t *testing.T, dir, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run %s %v: %v\nOutput:\n%s", name, args, err, string(output))
	}
}

// CreateTestGitRepo helper to create a temporary git repo
func CreateTestGitRepo(t *testing.T) string {
	baseDir := t.TempDir()

	run := func(name string, args ...string) {
		RunCmd(t, baseDir, name, args...)
	}

	// Initialize an empty repo
	run("git", "init")
	// Create an empty commit
	os.WriteFile(filepath.Join(baseDir, "README.md"), []byte("# test"), 0644)
	run("git", "add", ".")
	run("git", "config", "user.name", "test")
	run("git", "config", "user.email", "test@example.com")
	run("git", "config", "commit.gpgsign", "false")
	run("git", "commit", "-m", "init")

	return baseDir
}
