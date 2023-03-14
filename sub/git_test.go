package sub

import (
	"os/exec"
	"strings"
	"testing"
)

var (
	testHasGit bool
)

func init() {
	if _, err := exec.LookPath("git"); err == nil {
		testHasGit = true
	}
}

func TestGitStatus(t *testing.T) {
	// test skipped if git not found, one option
	if !testHasGit {
		t.Skip("Git command not found")
	}
	want := "On branch master"
	got, err := GitStatus()
	if err != nil {
		t.Fatalf("Gitstatus() err = %s; want nil", err)
	}
	if !strings.Contains(got, want) {
		t.Errorf("GitStatus() = %s; want prefix %s", got, want)
	}
}

func TestGitStatusMockCommand(t *testing.T) {
	// if our machine does not have git we are mocking an exec cmd to echo on branch master
	if !testHasGit {
		execCommand = func(name string, arg ...string) *exec.Cmd {
			return exec.Command("echo", "On branch master")
		}
		defer func() {
			execCommand = exec.Command
		}()
	}
	want := "On branch master"
	got, err := GitStatus()
	if err != nil {
		t.Fatalf("Gitstatus() err = %s; want nil", err)
	}
	if !strings.Contains(got, want) {
		t.Errorf("GitStatus() = %s; want prefix %s", got, want)
	}
}
