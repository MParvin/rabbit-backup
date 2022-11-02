package tools

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

// Git is a struct to hold git commands
type Git struct {
	// GitPath is the path to git executable
	GitPath string
}

// NewGit creates a new Git struct
func NewGit() *Git {
	return &Git{
		GitPath: "/usr/bin/git",
	}
}

// Clone clones a repository
func (g *Git) Clone(url, dir string) error {
	// git clone config.git.url config.git.dir
	cmd := exec.Command(g.GitPath, "clone", url, dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Add adds a file to git
func (g *Git) Add(file string) error {
	// git add config.git.dir
	cmd := exec.Command(g.GitPath, "add", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Commit commits a file to git
func (g *Git) Commit(file, message string) error {
	// git commit -m "Backup 2021-08-01 12:00:00" config.git.dir
	cmd := exec.Command(g.GitPath, "commit", "-m", message, file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Push pushes a file to git
func (g *Git) Push(file string) error {
	// git push config.git.dir
	cmd := exec.Command(g.GitPath, "push", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Checkout creates a new branch and switches to it
func (g *Git) Checkout(branch string) error {
	// git checkout -b config.git.branch
	cmd := exec.Command(g.GitPath, "checkout", "-b", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// BranchExists checks if a branch exists
func (g *Git) BranchExists(branch string) (bool, error) {
	// git branch -a
	cmd := exec.Command(g.GitPath, "branch", "-a")
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	// Check if branch exists
	if strings.Contains(string(out), branch) {
		return true, nil
	}
	return false, nil
}

// CheckoutBranch switches to an existing branch
func (g *Git) CheckoutBranch(branch string) error {
	// git checkout config.git.branch
	cmd := exec.Command(g.GitPath, "checkout", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Pull pulls a file from git
func (g *Git) Pull(file string) error {
	// git pull config.git.dir
	cmd := exec.Command(g.GitPath, "pull", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// GetBranch returns the current branch
func (g *Git) GetBranch() (string, error) {
	// git branch --show-current
	cmd := exec.Command(g.GitPath, "branch", "--show-current")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// GetRemote returns the current remote
func (g *Git) GetRemote() (string, error) {
	// git remote get-url origin
	cmd := exec.Command(g.GitPath, "remote", "get-url", "origin")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func GitDirExists(dir string) bool {
	// Check if git directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func GitDirEmpty(dir string) bool {
	// Check if git directory is empty
	f, err := os.Open(dir)
	if err != nil {
		return false
	}
	defer f.Close()
	_, err = f.Readdir(1)
	return err == io.EOF
}

func DeleteDir(dir string) error {
	// Delete directory
	return os.RemoveAll(dir)
}

func ChangeDir(dir string) error {
	// Change directory
	return os.Chdir(dir)
}
