package collectors

import "os/exec"

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
