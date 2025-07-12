package collectors

import (
	"os/exec"
	"strings"
)

// isGitDirectory check if `path` is a git repository.
func isGitDirectory(path string) bool {
		cmd := exec.Command("git", "-C", path, "rev-parse", "--is-inside-work-tree")
		err := cmd.Run()
		return err == nil
}

// isGitDirty test if the given git repository is dirty (has
// uncommitted changes).
func isGitDirty(path string) bool {
		cmd := exec.Command("git", "-C", path, "status", "--porcelain")
		output, err := cmd.Output()
		if err != nil {
				return false
		}

		return len(output) > 0
}

// getGitBranch get the current branch name of the git repostiroy.
func getGitBranch(path string) string {
		cmd := exec.Command("git", "-C", path, "rev-parse", "--abbrev-ref", "HEAD")
		output, err := cmd.Output()
		if err != nil {
				return ""
		}

		branch := strings.TrimSpace(string(output))
		return branch
}
