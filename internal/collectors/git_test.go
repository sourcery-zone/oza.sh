package collectors

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// runCmd run a command line in a directory and fail if error happened.
func runCmd(t *testing.T, dir ,name string, args ...string) {
		cmd := exec.Command(name, args...)
		cmd.Dir = dir
		output, err := cmd.CombinedOutput()
		if err != nil {
				t.Fatalf("Failed to run %s %v: %v\nOutput:\n%s", name, args, err, string(output))
		}
}

// createTestGitRepo helper to create a temporary git repo
func createTestGitRepo(t *testing.T) string {
		baseDir := t.TempDir()

		run := func(name string, args ...string) {
				runCmd(t, baseDir, name, args...)
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
		if !isGitDirectory(createTestGitRepo(t)) {
				t.Errorf("Expected directory to be a git repo")
		}

		// Test a normal directory
		if isGitDirectory(t.TempDir()) {
				t.Errorf("Expected to not be a git repository")
		}
}

// TestIsGitDirty test if the git repository is dirty
func TestIsGitDirty(t *testing.T) {
		repo := createTestGitRepo(t)

		if isGitDirty(repo) {
				t.Errorf("Expected the git repository to be clean!")
		}

		// Create an empty commit
		os.WriteFile(filepath.Join(repo, "LICENSE.md"), []byte("# test"), 0644)

		if !isGitDirty(repo) {
				t.Errorf("Expected the repo to be empty!")
		}

}
