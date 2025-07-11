package collectors

import "os/exec"

// isGitDirectory check if `path` is a git repository.
func isGitDirectory(path string) bool {
		cmd := exec.Command("git", "-C", path, "rev-parse", "--is-inside-work-tree")
		err := cmd.Run()
		return err == nil
}
