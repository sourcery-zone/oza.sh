package collectors

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sourcery-zone/oza.sh/internal/helpers"
)

// TestIsGitDirectory test if the input directory is a git repository.
func TestIsGitDirectory(t *testing.T) {
	if !isGitDirectory(helpers.CreateTestGitRepo(t)) {
		t.Errorf("Expected directory to be a git repo")
	}

	// Test a normal directory
	if isGitDirectory(t.TempDir()) {
		t.Errorf("Expected to not be a git repository")
	}
}

// TestIsGitDirty test if the git repository is dirty
func TestIsGitDirty(t *testing.T) {
	repo := helpers.CreateTestGitRepo(t)

	if isGitDirty(repo) {
		t.Errorf("Expected the git repository to be clean!")
	}

	// Create an empty commit
	os.WriteFile(filepath.Join(repo, "LICENSE.md"), []byte("# test"), 0644)

	if !isGitDirty(repo) {
		t.Errorf("Expected the repo to be empty!")
	}

}

// TestGetGitBranch get the current branch name of the git repository.
func TestGetGitBranch(t *testing.T) {
	repo := helpers.CreateTestGitRepo(t)

	if branch := getGitBranch(repo); branch != "main" && branch != "master" {
		t.Errorf("Expected to receive the main branch, received: %s", branch)
	}

	expect := "newBranch"
	helpers.RunCmd(t, repo, "git", "checkout", "-B", expect)

	if got := getGitBranch(repo); got != expect {
		t.Errorf("Expected: %s, got: %s", expect, got)
	}
}
