package collectors

import (
	"os/exec"
	"strings"
	"sync"
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

type GitStatus struct {
	IsGit  bool
	Branch string
	Dirty  bool
}

type Git struct {
	path string
	once sync.Once
	data GitStatus
}

// NewGit create a new `Git` object initialized with path.
func NewGit(path string) *Git {
	if path == "" {
		panic("Git: path is required")
	}
	return &Git{path: path}
}

// load the git repository information.
func (g *Git) load() {
	g.once.Do(func() {
		g.data = GitStatus{
			IsGit: isGitDirectory(g.path),
			Branch: getGitBranch(g.path),
			Dirty: isGitDirty(g.path),
		}
	})
}

// IsGit check if the current path is a git repository.
func (g *Git) IsGit() bool {
	g.load()
	return g.data.IsGit
}

// Branch get current branch name in the repository.
func (g *Git) Branch() string {
	g.load()
	return g.data.Branch
}

// Dirty get if the repository is dirty (had uncommitted change).
func (g *Git) Dirty() bool {
	g.load()
	return g.data.Dirty
}
