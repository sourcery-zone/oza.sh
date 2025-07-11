package collectors

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// createTestGitRepo helper to create a temporary git repo
func createTestGitRepo(t *testing.T) string {
		baseDir := t.TempDir()

		run := func(name string, args ...string) {
				cmd := exec.Command(name, args...)
				cmd.Dir = baseDir
				output, err := cmd.CombinedOutput()
				if err != nil {
						t.Fatalf("Failed to run %s %v: %v\nOutput:\n%s", name, args, err, string(output))
				}
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

// TestIsGitDirectory test if the input directory is a git repository.
func TestIsGitDirectory(t *testing.T) {
		dir := createTestGitRepo(t)

		if !isGitDirectory(dir) {
				t.Errorf("Expected directory to be a git repo")
		}
}

// TestIsNotGitDirectory ensure it's also able to check non-git directories.
func TestIsNotGitDirectory(t *testing.T) {
		dir := t.TempDir()

		if isGitDirectory(dir) {
				t.Errorf("Expected to not be a git repository")
		}
}
